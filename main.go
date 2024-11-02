package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gocolly/colly"
	// "github.com/gocolly/colly/proxy"
)

func main() {
	// Criar um novo coletor
	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"),
	)

	// // Configurar proxies
	// proxies := []string{
	// 	"http://proxy1.com",
	// 	"http://proxy2.com",
	// }

	// rp, err := proxy.RoundRobinProxySwitcher(proxies...)
	// if err != nil {
	// 	log.Fatalln("Falha ao configurar proxies", err)
	// }

	// c.SetProxyFunc(rp)

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

	// Handle de erro
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Erro ao acessar", r.Request.URL, ":", err)
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
		// Atraso aleatório entre requisições
		delay := time.Duration(rand.Intn(10)) * time.Second
		time.Sleep(delay)
	})

	// Extrair preço
	c.OnHTML(".price-class", func(e *colly.HTMLElement) {
		price := e.Text
		fmt.Println("Preço encontrado:", price)
		writer.Write([]string{price})
	})

	// Mensagem ao terminar
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Scraping finalizado:", r.Request.URL)
	})

	// Iniciar scraping no site desejado com timeout
	fmt.Println("Iniciando scraping em: https://example.com")
	err = c.Visit("https://example.com")
	if err != nil {
		log.Fatalln("Falha ao iniciar scraping", err)
	}
}
