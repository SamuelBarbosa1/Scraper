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

## Atualizações 
- 1. Captura de sinais do sistema (SIGINT e SIGTERM):
Para detectar quando Ctrl + C é pressionado e permitir que o programa execute tarefas de encerramento (como gerar o relatório em PDF e enviar o e-mail) antes de sair.

- 2. Função para geração de relatório PDF (gerarRelatorioPDF):
Essa função cria um relatório PDF dos preços coletados durante o scraping, salvando-o localmente como relatorio.pdf.

- 3. Envio de e-mail de notificação (enviarNotificacao):
    A função envia um e-mail informando que o processo de scraping foi concluído com sucesso.
    Utiliza o pacote gomail para enviar e-mails via SMTP.

- 4. Manuseio de erros de rede:
    Tratamento de erros de conexão com OnError e uma lógica para retry automático em caso de falhas (exemplo: r.Request.Retry()).

- 5. Atraso aleatório entre requisições:
    Adicionamos um atraso de tempo aleatório entre requisições para evitar sobrecarregar o site que está sendo raspado.

- 6. Goroutine para capturar sinais e realizar tarefas de limpeza:
    Uma goroutine foi configurada para aguardar um sinal de interrupção e garantir que as tarefas de geração de PDF e envio de e-mail sejam realizadas antes do encerramento do programa.