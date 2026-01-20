package main

import "fmt"

func main() {

	fmt.Println("Informe o sobrenome da sua Familia: ")
	var familia string
	fmt.Scanf("%s", &familia)

	if familia == "calvo" {
		fmt.Println("Integrante da familia Calvo.")
	} else {
		fmt.Println("Um bastardo qualquer.")
	}
}
