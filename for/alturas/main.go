package main

import "fmt"

func main() {
	fmt.Println("Calculo de Alturas")
	var alturas float32
	var sum float32
	for i := 0; i < 4; i++ {
		fmt.Printf("Informe a %d  altura: ", i+1)

		fmt.Scanf("%f", &alturas)
		sum += alturas
	}
	fmt.Printf("a soma das alturas é : %.2f \n", sum)
}
