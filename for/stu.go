package main

import "fmt"

type camisa struct {
	cor     string
	tamanho float64
	gola    string
	manga   bool
}

func main() {
	fmt.Println("hhh")
	camisa_botafogo := camisa{"preta", 40.2, "v", true}

	fmt.Println(camisa_botafogo)

}
