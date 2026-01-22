package main

import "fmt"

func main() {
	var sum, entrada float32
	for i := 0; ; i++ {

		fmt.Println("Informe o valor desejado: ")
		fmt.Scanf("%f", &entrada)
		if entrada == 5 {
			break

		}
		sum += entrada

	}
	fmt.Printf("%.2f\n", sum)

}
