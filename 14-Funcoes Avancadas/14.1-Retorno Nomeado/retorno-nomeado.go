package main

import "fmt"


/*
	No retorno nomeado não é mais necessário declarar as variáveis dentro do bloco da função
	você declara junto com a função, e no final de seu bloco é só chamar o return
	mas cuidado, a ordem dos fatores interfere.

	os retornos nomeados devem ser declarados dentro de parenteses, mesmo se for um só
*/


func calculosMatematicos(n1, n2 int)(soma int, subtracao int){
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main(){
	fmt.Println(calculosMatematicos(1,2))
}