# Steam Scraper

Meu primeiro projeto em **Go**, desenvolvido para minerar dados da Steam e encontrar as melhores promoções com base no orçamento do usuário.

## Funcionalidades
- **Web Scraping:** Extração de dados diretamente do HTML da Steam usando `goquery`.
- **Filtro de Preço:** Conversão de strings complexas (ex: "R$ 45,50") para `float64` para filtragem lógica.
- **Detecção de Descontos:** Identifica a porcentagem de desconto e limpa a formatação do site.
- **User-Agent Spoofing:** Simulação de navegador para evitar bloqueios de robôs.

## Tecnologias
- **Linguagem:** Go (Golang)
- **Biblioteca Principal:** [Goquery](https://github.com/PuerkitoBio/goquery)

## Como rodar
1. Certifique-se de ter o Go instalado.
2. Clone o repositório.
3. Instale as dependências:
   ```bash
   go mod tidy
4. Execute o programa:
   ```bash
   go run main.go
