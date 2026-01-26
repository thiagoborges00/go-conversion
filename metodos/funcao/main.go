package main

import "fmt"

func media(a, b float32) (result float32) {
	result = (a + b) / 2
	return
}

func main() {
	fmt.Println("Informe dois numeros e calcularemos a media deles: ")
	var a, b, med float32
	fmt.Scanf("%f", &a)
	fmt.Scanf("%f", &b)

	med = media(a, b)

	fmt.Printf("O valor da media e: %.2f\n", med)
}
