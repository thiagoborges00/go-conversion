package main

import "fmt"

func main() {
	fmt.Println("Bom dia, Informe o seu nome")
	var nome string
	fmt.Scanf("%s", &nome)

	fmt.Printf("prazer rm conhecer %s , seja bem vindo(a)\n", nome)
}
