package main

import (
	"errors"
	"fmt"
)

func main() {
	//INTEIROS
	//	int8, int16, int32, int64  (quanto maior o número mais digitos ele aceita, positivo ou negativo)
	var numero int64 = 1000000000000000000
	numero2 := 1000000000000000000 //se não delcarar o tipo o go usa a arquitetura do seu pc 64 ou 32, no meu caso é 64
	fmt.Println(numero)
	fmt.Println(numero2)
	var numero3 uint32 = 2000 //unsygned int  (não pode ser negativo)
	fmt.Println(numero3)

	//alias
	var numero5 rune = 1256 //igual ao int32 este rune
	var numero6 byte = 123  //byte igual ao int8
	fmt.Println(numero5, numero6)

	//FLOAT
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 5151515151.56
	fmt.Println(numeroReal2)

	numeroReal3 := 12653.56 //de acordo com a arquitetura do procesador, no meu caso será um float64
	fmt.Println(numeroReal3)

	//STRING
	var str string = "Texto" // sempre aspas duplas
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B'       //inferiu o tipo RUNE, uma ínica string declarada com aspas simples vira o número da tabela ascii que representa esse caracter
	fmt.Println(char) //mostra o numero da tabela ascii

	//VALOR 'ZERO'
	//toda variável tem seu valor zero se você declarar e não passar um valor, string é string vazia e numeros são 0
	var texto float64
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var booleano2 bool //false é o valor ZERO
	fmt.Println(booleano2)

	var erro error
	fmt.Println(erro) //<nil> é o valor ZERO

	var erro2 error = errors.New("Erro Interno")
	fmt.Println(erro2)

}
