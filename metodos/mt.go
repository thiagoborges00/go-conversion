package main

import "fmt"

type aretangulo struct {
	comprimento float64
	altura      float64
}

func (casa *aretangulo) area() float64 {
	return casa.altura * casa.comprimento
}

// método é uma função associada a um tipo, só funciona
// para o tipo que foi declarado

func main() {
	a := aretangulo{8.0, 42}
	fmt.Println(a.area())

}
