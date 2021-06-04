package main

import "fmt"



func main(){
	i := 0
	for i < 10 {
		i++
		fmt.Println(i)
	}
	fmt.Println(i)




	for j:=0; j<=10; j+=2{
		fmt.Println("Incrementando J", j)
	}


	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, valor := range nomes {
		fmt.Println(indice, valor)
	}

	for _, valor := range nomes {
		fmt.Println(valor)
	}

	for indice, letra := range "PALAVRA"{
		fmt.Println(indice, string(letra))  //se nao fizer assim é retornado o numero da tabela ascii
	}

	//RANGE EM MAPS
	usuario := map[string]string{
		"nome": "Fernando",
		"idade": "29",
	}


	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//NÃO TEM COMO FAZER RANGE EM STRUCT


	//LOOP INFINITO
	// for {
	// 	fmt.Println("Executando infinitamente")
	// }





}