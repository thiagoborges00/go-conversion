package main

import "fmt"

func main() {

	fmt.Println("Informe o número que deseja dobrar: ")
	var num int
	fmt.Scanf("%d", &num)
	fmt.Printf("O dobro do numero desejado é: %d\n", num*2)
}
