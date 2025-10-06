package main

import "fmt"

func main() {
	fmt.Println("milk")

	// fatia sem variavel declarada
	fatia := make([]float64, 4)

	// com a variavel declarada ela precisa ser inicializada
	array_fatia2 := [4]float64{1, 2, 3, 4}

	fmt.Println(fatia)
	fmt.Println(array_fatia2[2:4])

	// acrescentando elemento na fatia

	fatia3 := append(array_fatia2[:], 700, 800)
	fmt.Println(fatia3)

	//copiar uma fatia na outra
	seriea := make([]string, 3)
	seriea = append(seriea, "fortaleza", "mirasol", "fluminense")
	serieb := make([]string, 2)
	serieb = append(serieb, "paysandu", "atlhetico")
	copy(seriea, serieb)
	fmt.Println(seriea, serieb)
}
