package main

import "fmt"

func main() {
	x := "vitoria"
	if x == "santos" {
		fmt.Println("Santos")
	} else if x == "flamengo" {
		fmt.Println("Flamengo")
	} else {
		fmt.Println("Corinthians")
	}
}
