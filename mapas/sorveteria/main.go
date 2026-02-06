package main

import "fmt"

func main() {

	println("Sorveteria Di Gelat")
	tipoSorvete := map[string]float32{
		"casquinha": 1.0,
		"cascao":    2.0,
		"cestinha":  4.0,
	}

	saborSorvete := map[string]float32{
		"morango":   0.0,
		"creme":     0.0,
		"chocolate": 0.0,
	}

	saborCobertura := map[string]float32{
		"caramelo":  1.5,
		"morango":   1.5,
		"chocolate": 1.5,
		"sem":       0,
	}

	itens := map[string]map[string]float32{}
	var sabor, cobertura, tipo string
	var valor float32

	fmt.Printf("Escolha o tipo de sorvete [ casquinha, cascao, cestinha ]")
	fmt.Scanf("%s", &tipo)

	fmt.Printf("Escolha sabor do sorvete: [ morango, creme, chocolate ]")
	fmt.Scanf("%s", &sabor)

	fmt.Printf("Escolha o sabor da cobertura: [ morango, chocolate, caramelo, sem cobertura ]")
	fmt.Scanf("%s", &cobertura)

	itens["tipo"] = tipoSorvete
	itens["sabor"] = saborSorvete
	itens["cobertura"] = saborCobertura

	if v1, ok := itens["tipo"][tipo]; ok {
		valor += v1
	} else {
		fmt.Println("Tipo de sorvete inválido!")
		return
	}

	if v2, ok := itens["sabor"][sabor]; ok {
		valor += v2
	} else {
		fmt.Println("Sabor não encontrado!")
		return
	}

	if v3, ok := itens["cobertura"][cobertura]; ok {
		valor += v3
	} else {
		fmt.Println("Cobertura não encontrada!")
		return
	}

	fmt.Printf("O total do Pedido é de: R$ %.2f\n", valor)

}
