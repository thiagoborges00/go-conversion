package main

import "fmt"

func main() {
	fmt.Println("Barraca de venda de H2O")
	fmt.Println("qual agua voce deseja: 1 - natural, 2 - com gas:")
	var entrada string
	fmt.Scanf("%s", &entrada)
	fmt.Println("Quantas unidades deseja?")
	var qtd int
	fmt.Scanf("%d", &qtd)
	preco := 0.0

	if entrada == "1" || entrada == "natural" {

		preco += 1.5
	} else {
		preco += 2.5

	}
	fmt.Printf("O valor da sua compra é de R$ %.2f\n", preco*float64(qtd))
}
