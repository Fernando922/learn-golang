package main

import "fmt"


func main(){
	//ARITMETICOS

	soma := 1+2
	subtracao := 1-2
	divisao := 10/4
	restoDaDivisao := 10 % 2


	fmt.Println(soma, subtracao, divisao, restoDaDivisao)

	//Uma variável de um tipo não pode interagir com outra de outra tipo
	// var numero1 int16 = 10
	// var numero2 int32 = 25
	// soma := numero1 + int16(numero2)  //não pode porque são tipos diferentes int16 e int32



	//OPERADORES DE ATRIBUIÇÃO
	var variavel string = "String"  //declaração tipada
	variavel2 := "String2"  //inferência de tipo

	fmt.Println(variavel, variavel2)


	//OPERADORES RELACIONAIS
	fmt.Println(1>2)
	fmt.Println(1>=2)
	fmt.Println(1==2)
	fmt.Println(1<2)
	fmt.Println(1<=2)
	fmt.Println(1!= 2)


	//fmt.Println(1=="1")  //vai dar erro (ainda bem) porque você está comparando tipos diferentes


	fmt.Println("----------")


	//OPERADORES LOGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)


	//OPERADORES UNÁRIOS
	numero :=10 
	numero++   //Não existe --numero ou ++numero
	fmt.Println(numero)
	numero+=15
	fmt.Println(numero)
	numero--
	numero-=20
	fmt.Println(numero)
	numero*=5
	fmt.Println(numero)
	numero/=10
	fmt.Println(numero)
	numero%=2
	fmt.Println(numero)

	//OPERADOR TERNÁRIO
	//texto := numero > 5 ? "Maior que 5": "Menor que 5"  não existe em GO

	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	}else{
		texto= "Menor que 5"
	}
	println(texto)








}