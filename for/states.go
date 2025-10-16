package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			fmt.Println("Bahia")
		case 1:
			fmt.Println("Ceará")
		case 2:
		case 3, 4:
			fmt.Println("Pernanbuco")
		case 5:
			fmt.Println("Maranhão")
			fmt.Println("Piauí")

		}
	}

}
