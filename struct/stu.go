package main

import "fmt"

type camisa struct {
	cor     string
	tamanho float64
	gola    string
	manga   bool
	suja    bool
}

/*
para criar funções ou metodos que alterem a struct
a assinatura fica no modelo   func (cm *camisa) PassarCamisa(){}

de forma tal que o ponteiro da estrutura garante que sempre será alterado
a instancia original
*/

func (cm *camisa) LavarCamisa() {
	cm.suja = true
}

//tipos incluidos ou embedding
// acontece quando ocorre algo parecido com a heranca

type camisaSocial struct {
	camisa
	colarinho string
	modelo    string
}

func main() {
	fmt.Println("iniciando")
	camisa_botafogo := camisa{"preta", 40.2, "v", true, false}
	camisa_bahia := camisa{}
	camisa_bahia.cor = "azul"

	camisaItaliana := camisaSocial{
		camisa: camisa{
			cor:     "vermelha",
			tamanho: 21,
			gola:    "v",
			manga:   true,
			suja:    false,
		},
		colarinho: "largo",
		modelo:    "italiano",
	}

	fmt.Println(camisa_botafogo.gola)
	fmt.Println(camisa_bahia)
	fmt.Println(camisaItaliana)

}
