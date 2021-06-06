package main

import "fmt"

//São funções que referenciam variáveis fora de seu corpo
//A função mantem o conteudo da variável que foi passado para ela no momento da criação


func closure() func() {
	texto := "Dentro da função closure"


	funcao := func(){
		fmt.Println(texto)
	}

	return funcao
}

func main(){
	texto := "Dentro da função main"
	fmt.Println(texto)


	funcaoNova := closure()
	funcaoNova()  //Dentro da função closure
}