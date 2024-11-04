package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gocolly/colly"
	"github.com/jung-kurt/gofpdf"
	"gopkg.in/gomail.v2"
)

func formatPrice(rawPrice string) string {
	formatted := strings.TrimSpace(rawPrice)
	formatted = strings.ReplaceAll(formatted, "$", "")
	formatted = strings.ReplaceAll(formatted, ",", "")
	return formatted
}

func gerarRelatorioPDF(precos []string) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "Relatório de Preços")

	for _, preco := range precos {
		pdf.Ln(10)
		pdf.Cell(0, 10, preco)
	}

	err := pdf.OutputFileAndClose("relatorio.pdf")
	if err != nil {
		log.Fatalln("Erro ao criar PDF:", err)
	}
}

func enviarNotificacao() {
	m := gomail.NewMessage()
	m.SetHeader("From", "seuemail@exemplo.com")
	m.SetHeader("To", "destinatario@exemplo.com")
	m.SetHeader("Subject", "Scraping Concluído")
	m.SetBody("text/plain", "O scraping foi concluído com sucesso!")

	d := gomail.NewDialer("smtp.gmail.com", 587, "seuemail@exemplo.com", "senha")
	if err := d.DialAndSend(m); err != nil {
		log.Fatalln("Erro ao enviar e-mail:", err)
	}
}

func main() {
	// Canal para capturar sinais de interrupção
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Criar um novo coletor
	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"),
	)

	// Limitar a profundidade de busca e paralelismo
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		Delay:       2 * time.Second,
	})

	// Criar arquivo CSV para salvar dados
	file, err := os.Create("dados.csv")
	if err != nil {
		log.Fatalln("Falha ao criar arquivo", err)
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()

	var precos []string

	// Handle de erro
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Erro ao acessar", r.Request.URL, ":", err)
		if r.StatusCode != 0 {
			fmt.Println("Tentando novamente:", r.Request.URL)
			time.Sleep(5 * time.Second)
			r.Request.Retry()
		}
	})

	// Log de requisição
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visitando", r.URL.String())
	})

	// Encontrar e visitar todos os links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println("Link encontrado:", link)
		c.Visit(e.Request.AbsoluteURL(link))
		delay := time.Duration(rand.Intn(10)) * time.Second
		time.Sleep(delay)
	})

	// Extrair preço
	c.OnHTML(".price-class", func(e *colly.HTMLElement) {
		price := formatPrice(e.Text)
		fmt.Println("Preço formatado encontrado:", price)
		writer.Write([]string{price})
		precos = append(precos, price)
	})

	// Mensagem ao terminar
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Scraping finalizado:", r.Request.URL)
		enviarNotificacao()
		gerarRelatorioPDF(precos)
	})

	// Goroutine para capturar sinal e realizar ações de limpeza
	go func() {
		<-sigs
		fmt.Println("\nInterrupção detectada. Finalizando e gerando relatórios...")
		gerarRelatorioPDF(precos)
		enviarNotificacao()
		os.Exit(0)
	}()

	// Iniciar scraping no site desejado com timeout
	fmt.Println("Iniciando scraping em: https://example.com")
	err = c.Visit("https://example.com")
	if err != nil {
		log.Fatalln("Falha ao iniciar scraping", err)
	}
}

//eee
