package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gocolly/colly"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func formatPrice(rawPrice string) string {
	formatted := strings.TrimSpace(rawPrice)
	formatted = strings.ReplaceAll(formatted, "$", "")
	formatted = strings.ReplaceAll(formatted, ",", "")
	return formatted
}

func enviarSMS(mensagem string) {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: "YOUR_TWILIO_ACCOUNT_SID",
		Password: "YOUR_TWILIO_AUTH_TOKEN",
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo("+5561999544335")             // Número do destinatário
	params.SetFrom("YOUR_TWILIO_PHONE_NUMBER") // Número Twilio registrado
	params.SetBody(mensagem)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Println("Erro ao enviar SMS:", err)
	} else {
		fmt.Println("SMS enviado com sucesso:", *resp.Sid)
	}
}

func main() {
	// Configuração de sinal para capturar interrupções
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Criar um novo coletor
	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, como Gecko) Chrome/58.0.3029.110 Safari/537.3"),
	)

	// Configurar limites
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		Delay:       2 * time.Second,
	})

	var precos []string

	// Configurar handlers do coletor
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Erro ao acessar", r.Request.URL, ":", err)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visitando", r.URL.String())
	})

	c.OnHTML(".price-class", func(e *colly.HTMLElement) {
		price := formatPrice(e.Text)
		fmt.Println("Preço encontrado:", price)
		precos = append(precos, price)
	})

	// Goroutine para capturar sinal de interrupção
	go func() {
		<-sigs
		fmt.Println("\nInterrupção detectada. Finalizando...")
		mensagem := "Relatório de Preços: " + strings.Join(precos, ", ")
		enviarSMS(mensagem)
		os.Exit(0)
	}()

	// Iniciar o scraping
	fmt.Println("Iniciando scraping em: https://example.com")
	err := c.Visit("https://example.com")
	if err != nil {
		log.Fatalln("Falha ao iniciar scraping", err)
	}

	// Mantém o programa rodando
	select {}
}
