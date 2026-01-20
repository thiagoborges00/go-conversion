package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Informe o número que deseja extrair a raiz quadrada: ")
	var sqrt float64
	fmt.Scanf("%f", &sqrt)
	fmt.Printf("A raiz quadrada de %.2f é:  %.2f:", sqrt, math.Sqrt(sqrt))
}
