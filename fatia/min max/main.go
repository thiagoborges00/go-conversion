package main

import "fmt"

func main() {
	var notas [4]float32
	var min float32 = 100.0
	var max, sum, nota float32
	for i := 0; i < 4; i++ {
		fmt.Printf("Informe a %d nota do aluno : ", i+1)
		fmt.Scanf("%f", &nota)
		notas[i] = nota
		if notas[i] > max {
			max = notas[i]

		}
		if notas[i] < min {
			min = notas[i]
		}

		sum += notas[i]
	}
	fmt.Printf("Min: %.2f\nMax: %.2f\nMedia: %.2f\n", min, max, sum/float32(len(notas)))
}
