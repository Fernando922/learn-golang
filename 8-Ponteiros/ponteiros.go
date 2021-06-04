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
	fmt.Println(variavel1, variavel2)  //11  10  só a primeira foi alterada


	//PONTEIRO É UMA REFERENCIA DE MEMORIA
	//VOCE ATRIBUI A REFERENCIA DE MEMÓRIA ONDE A VARIÁVEL ESTÁ SALVA
	//ENTÃO SE UM VALOR ALTERAR, NA OUTRA VARIÁVEL TB VAI ALTERAR

	var variavel3 int  //0 
	var ponteiro *int  //<nil>


	variavel3 = 100
	ponteiro = &variavel3  //adiciona com E comercial na frente do nome


	fmt.Println(variavel3, ponteiro)  //100 0xc0000b6038 (endereço de memória)
	fmt.Println(variavel3, *ponteiro)  //desreferenciação você adicionou o asterisco para ler o valor atribuido

	variavel3 = 150
	fmt.Println(variavel3, ponteiro)  //150 0xc0000b6038 (endereço de memória sempre vai manter)
	fmt.Println(variavel3, *ponteiro)  //150 150




}