package main

import "fmt"

func main() {

	var notas []float32 // dessa forma não aloca , fica com tamanho dinamico

	var entrada, soma float32
	fmt.Println("Informe as notas do aluno")

	for i := 0; ; i++ {
		fmt.Printf("Informe a %d nota, digite 100 para sair. ", i+1)
		fmt.Scanf("%f", &entrada)

		if entrada == 100 {
			break
		}

		notas = append(notas, entrada)
		soma += entrada
	}

	fmt.Printf("notas: %v, soma %.2f\n", notas, soma)
}
