package main

import (
	"fmt"

)

//variavel que salva endereço de memória de alguma coisa



func main(){
	var variavel1 int = 10
	var variavel2 int = variavel1  //copia do valor da primeira


	fmt.Println(variavel1, variavel2)  //10   10


	variavel1++
	fmt.Println(variavel1, variavel2)  //11  10  só a primeira foi alterada, porque a segunda foi uma cópia do valor da primeira


	//PONTEIRO É UMA REFERENCIA DE MEMORIA
	//VOCE ATRIBUI A REFERENCIA DE MEMÓRIA ONDE A VARIÁVEL ESTÁ SALVA
	//ENTÃO SE UM VALOR ALTERAR, NA OUTRA VARIÁVEL TB VAI ALTERAR

	var variavel3 int  //0 

	//para tipar um novo ponteiro, ele deve deve ter o *antes do tipo
	var ponteiro *int  //<nil> é seu valor zero


	variavel3 = 100
	ponteiro = &variavel3  //adiciona com E comercial na frente do nome para pegar a referencia de memória e não seu valor
	var ponteiro2 *int = &variavel3
	ponteiro3 := &variavel3

	minhaString := "Fernando"
	ponteiro4 := &minhaString
	minhaString = "Fernando alterado"


	fmt.Println(variavel3, ponteiro)  //100 0xc0000b6038 (endereço de memória)
	fmt.Println(variavel3, *ponteiro)  //desreferenciação você adicionou o asterisco para ler o valor atribuido

	variavel3 = 150
	fmt.Println(variavel3, ponteiro)  //150 0xc0000b6038 (endereço de memória sempre será exibido sem o asterisco vai manter)
	fmt.Println(variavel3, *ponteiro, *ponteiro2, *ponteiro3, *ponteiro4)  //150 150 150




}