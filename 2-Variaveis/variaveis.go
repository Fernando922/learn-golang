package main

import "fmt"


//go não compíla se tiver variavel declarada que não esteja sendo utilizada

func main() {
	var variavel1 string = "Variável 1"  //declaração tipada
	variavel2 := "Variável 2"  //inferencia de tipo, nao declara como 'var"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	//declarando multiplas variáveis
	var (
		variavel3 string = "variavel3"
		variavel4 string = "variavel4"
	)

	fmt.Println(variavel3, variavel4)

	//declarando múltiplas do mesmo tipo com inferencia
	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	//constante deve ser declarada como "const"
	const constante1 string = "Constante 1"
	fmt.Println(constante1)


	//por conta de existir multiplas declarações você pode trocar de valores sem variável auxiliar
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}