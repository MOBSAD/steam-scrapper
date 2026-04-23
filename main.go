package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	var url string = "https://store.steampowered.com/search"
	//fmt.Println("Insira o link da Steam:") //pode ser alterado depois
	//fmt.Scan(&url)

	var maxPriceStr string
	fmt.Println("Qual o teto que pode gastar? (Ex: 50.00)")
	fmt.Scan(&maxPriceStr)

	// converte o teto para float64 pra calcular
	// se não converter pipoca tudo kkkk
	teto, err := strconv.ParseFloat(maxPriceStr, 64)
	if err != nil {
		log.Fatal("Digite um valor numérico válido para o preço!")
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(res.Body)

	fmt.Printf("\n--- Jogos até R$ %.2f ---\n", teto)

	doc.Find(".search_result_row").Each(func(i int, s *goquery.Selection) {
		titulo := s.Find(".title").Text()
		precoTexto := s.Find(".discount_final_price").Text()
		desconto := s.Find(".discount_pct").Text()

		if desconto == "" {
			desconto = "0%"
		}
		// Se o preço estiver vazio skipa porque é perca de tempo SKIIIP
		if precoTexto == "" {
			return
		}

		// limpeza do valor dessa bosta, formatando a poha toda
		valorlimpo := strings.ReplaceAll(precoTexto, "R$", "")
		valorlimpo = strings.ReplaceAll(valorlimpo, ",", ".")
		valorlimpo = strings.TrimSpace(valorlimpo)

		// limpando o desconto pq sapoha da -[desconto]%
		descontolimpo := strings.ReplaceAll(desconto, "-", "")
		desconto = descontolimpo

		// converte valor do jogo pra float pra ficar no capricho :)
		precoJogo, err := strconv.ParseFloat(valorlimpo, 64)

		if err == nil {
			// compara sapoha, jogo gratis que se foda
			if precoJogo <= teto {
				fmt.Printf("Nome: %s\nPreço: R$ %.2f - [%s off]\n\n", titulo, precoJogo, desconto)
			}
		}
	})
}
