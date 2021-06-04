package main

//quantidade de parametros indefinida
//só pode ter um variatico por função e ele deve estar sempre por último

import "fmt"




func soma(numeros ...int)int{

	total := 0

	for _, numero := range numeros{

		total += numero

	}
	return total

}

func escrever(texto string, numeros...int){
	for _, numero :=range numeros {
		fmt.Println(texto, numero)
	}
	fmt.Println()
}


func main(){

	fmt.Println(soma(2,5,6,4,8,5))
	escrever("Fernando", 1,2,3,5)

}
