# Web Scraper em Go

Este é um projeto de Web Scraper simples usando a linguagem de programação Go e a biblioteca `colly, Twilio`. O scraper coleta dados de um site, salva em um relatorio e respeita limites e atrasos configuráveis.

## Funcionalidades

- Coleta de links e dados específicos (preços)
- Atrasos aleatórios entre requisições para evitar detecção
- Tratamento de erros e logs de requisições

## Como Usar

1. **Instale Go**: Certifique-se de que você tem Go instalado na sua máquina. [Download Go](https://golang.org/dl/)

2. **Clone o Repositório**:
    ```bashe
    git clone https://github.com/SamuelBarbosa1/Scraper.git
    cd Scraper
    ```

3. **Instale Dependências**:
    ```bash
    go mod init meu-projeto
    go get github.com/gocolly/colly
    go get github.com/gocolly/colly/proxy
    go get github.com/jung-kurt/gofpdf
    go get gopkg.in/gomail.v2
    ```

4. **Execute o Scraper**:
    ```bash
    go run main.go
    ```
## Exemplo 

```
go run main.go
Iniciando scraping em: https://books.toscrape.com/
Visitando https://books.toscrape.com/
Produto encontrado: A Light in the ... - Preço: 51.77
Produto encontrado: Tipping the Velvet - Preço: 53.74
Produto encontrado: Soumission - Preço: 50.10
Produto encontrado: Sharp Objects - Preço: 47.82
````
## Atualizações 
- 1. Captura de sinais do sistema (SIGINT e SIGTERM):
Para detectar quando Ctrl + C é pressionado e permitir que o programa execute tarefas de encerramento (como gerar o relatório e enviar pelo SMS) antes de sair.

- 2. Função para geração de relatório
Essa função cria um relatório  dos preços coletados durante o scraping, salvando-o localmente 

- 3. Removi a opção enviar pelo email, agora é apenas via SMS por questões de melhorias  de segurança.

- 4. Manuseio de erros de rede:
    Tratamento de erros de conexão com OnError e uma lógica para retry automático em caso de falhas (exemplo: r.Request.Retry()).

- 5. Atraso aleatório entre requisições:
    Adicionamos um atraso de tempo aleatório entre requisições para evitar sobrecarregar o site que está sendo raspado.

- 6. Removi o modo enviar por Email, pois estava dando muitos problemas por conta de autenticação ou coisa parecidas, então joguei tudo via SMS.

- 7. Adicionei um campo para o usuário informar o número do seu celular.

## Bugs 

- 1. Problema de autenticação no email, removi a opção de enviar.

- 2. Problema de atraso entre requisições, removi a opção.

- 3. Problema de tratamento de erros de rede, removi a opção.

## Sugestões de Melhoria 
- 1.  Implementar um sistema de cache para evitar requisições desnecessáreas ao remetente.

- 2.  Implementar um sistema de autenticação para evitar que qualquer pessoa possa enviar mensagens pelo sistema.

- 3.  Implementar um sistema de monitoramento para detectar problemas no sistema e enviar alertas para os administr.

- 4.  Implementar um sistema de backup para garantir a segurança dos dados do codigo.

- 5.  Implementar um sistema de atualização automática para garantir que o sistema esteja sempre atual.

- 6.  Implementar  um sistema de log para registrar todas as ações realizadas pelo sistema.

## Atualizações 

- 1. Removemos muitas linhas de codigos que não serão mais necessárias, assim podemos corrigir erros futuros com um codigo mais limpo.

- 2. Adicionamos um campo para o usuário informar o número do seu celular.

## Fotos
<p align="center"> <img src="./fotos/scrap.png" width="40"> <img src="./fotos/sms.png" width="40%"> </p>
