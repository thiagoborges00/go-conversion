package main

import "fmt"

func main() {
	fmt.Println("Informe uma palavra qualquer e vou lhe responder quantos 'a' tem nela: ")
	var palavra string
	fmt.Scanf("%s", &palavra)
	var count int8
	for _, value := range palavra {
		if value == 'a' {
			count++
		}
	}
	fmt.Printf("A quantidade total de 'a' é:  %d\n", count)
}
