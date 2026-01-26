package main

import "fmt"

// values funciona com o args do python,
// é uma fatia dinamica

// so pode ter um argumento variadico

func mediaComParametrosVariaveis(values ...float32) float32 {
	var soma float32 = 0
	for _, v := range values {
		soma += v
	}
	return soma / float32(len(values))
}

func main() {

	// atribuicao de fatias

	valores := []float32{45.5, 55.5, 200}
	// media := mediaComParametrosVariaveis(10, 32, 25.8, 40, 100)
	media := mediaComParametrosVariaveis(valores...)

	fmt.Println(media)
}
