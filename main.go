package main

import "fmt"

const tempK  float64  = 373.0

func main(){

	tempCelsius := tempK - 273 
	fmt.Println("A teperatura em celsius é : ", tempCelsius)
}