package main

import "fmt"

func main() {
	var notas [5]float64
	notas[0] = 7.7
	notas[1] = 7.9
	notas[2] = 6.2
	notas[3] = 8.1
	notas[4] = 9.4

	summ := 0.0
	for i := 0; i < len(notas); i++ {
		summ += notas[i]
	}
	fmt.Println(summ / float64(len(notas)))
}

// ALTERNATIVA usando for com range
// for index, value:= range notas{}

/*
declaracao de arrays
 - var nome[tamanho] tipo

inicializacao calculada
 - var precos := [4]float64{0: 1* 0.66, 1:6* 0.99, 2: 3*4.65}
 os indices omitidos recebem zero

array como argumento de funcao
na hora de enviar passar o endereço do array  calcula_media(&notas)

na hora de declarar a funcaao: func calcula_media(notas *[5]float64)
))
*/
