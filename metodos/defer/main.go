package main

import "fmt"
// defer: na pilha de execução garante que a funcao
// sera a ultima execucao



// panic  => como se fosse o exception do python

// trata erro de panic
defer func(){
	txt := recover()
	fmt.Println(txt)
}()

func main(){
	fmt.Printf("Defer")
}

