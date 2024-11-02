# Web Scraper em Go

Este é um projeto de Web Scraper simples usando a linguagem de programação Go e a biblioteca `colly`. O scraper coleta dados de um site, salva em um arquivo CSV e respeita limites e atrasos configuráveis.

## Funcionalidades

- Coleta de links e dados específicos (preços)
- Salva dados em um arquivo CSV
- Configuração de proxies para evitar bloqueios de IP
- Atrasos aleatórios entre requisições para evitar detecção
- Tratamento de erros e logs de requisições

## Como Usar

1. **Instale Go**: Certifique-se de que você tem Go instalado na sua máquina. [Download Go](https://golang.org/dl/)

2. **Clone o Repositório**:
    ```bash
    git clone https://github.com/SamuelBarbosa1/Scraper.git
    cd Scraper
    ```

3. **Instale Dependências**:
    ```bash
    go mod init meu-projeto
    go get github.com/gocolly/colly
    go get github.com/gocolly/colly/proxy
    ```

4. **Execute o Scraper**:
    ```bash
    go run main.go
    ```
## Exemplo 

```
go run main.go
Iniciando scraping em: https://example.com
Visitando https://example.com
Link encontrado: https://www.iana.org/domains/example
Visitando https://www.iana.org/domains/example
Link encontrado: /
Visitando http://www.iana.org/
Link encontrado: about/
Visitando http://www.iana.org/about/
```