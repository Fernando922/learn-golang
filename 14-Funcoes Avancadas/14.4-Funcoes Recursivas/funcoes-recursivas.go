package main

import "fmt"

/*
	Funções recursivas são função que dependem delas mesmas para executar, ou seja, é uma função
	que chama ela mesma até atingir sua condição de parada.
	Se não houver condição de parada acontecerá o que chamamos de estouro de pilha (stack overflow)
	que é onde cada função em execução é adicionado a uma pilha de execuções que tem um limite


	Uma função recursiva não é recomendada se for utilizar com inúmeras execuções, senão pode-se travar.

	Uma função recursiva é pouco utilizada na maioria das aplicações
*/

// 1  1  2  3  5  8  13   O próximo número é a soma dos anteriores

func fibonacci(posicao uint) uint { //uint é sempre maior que zero

	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)

}

func main() {

	posicao := uint(5)

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}
	fmt.Println(fibonacci(posicao))
}
