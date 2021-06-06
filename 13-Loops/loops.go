package main

import "fmt"

func main() {
	i := 0

	//for do mesmo estilo do WHILE sem definir o limite do contador e incrementar dentro do bloco
	for i < 10 {
		i++
		fmt.Println(i)
	}
	fmt.Println(i)

	//for com auto incremento já definido
	for j := 0; j <= 10; j += 2 {
		fmt.Println("Incrementando J", j)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}

	//iterando sobre cada item de um array
	for indice, valor := range nomes {
		fmt.Println(indice, valor)
	}

	//Não quero trabalhar com o indice do array então declaro com underline
	for _, valor := range nomes {
		fmt.Println(valor)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra)) //se nao fizer assim é retornado o numero da tabela ascii já que cada letra simples é uma representação da tabela
	}

	//RANGE EM MAPS
	usuario := map[string]string{
		"nome":  "Fernando",
		"idade": "29",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//NÃO TEM COMO FAZER RANGE EM STRUCT
	//PODE SE FAZER RANGE EM ARRAY MAP E SLICE

	//LOOP INFINITO
	// for {
	// 	fmt.Println("Executando infinitamente")
	// }

}
