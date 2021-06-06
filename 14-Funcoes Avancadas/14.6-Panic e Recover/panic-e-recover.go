package main

import "fmt"

/*
	Panic é chamado quando uma determinada condição for atingida
	siginifica que seu programa entrará em panico e chamará todas
	as funções DEFER declaradas.

	Panic é diferente de um erro, um erro você pode tratar e continuar a execução
	Panic MATA a execução do programa


	Panic sem um recover mata a execução de um programa

	Uma função de recover pode ser chamada tranquilamente mesmo se não houver panico
	uma mensagem de panico é recebida pelo recover() então é só tratar com if else

*/

func recuperarExecucao() {
	fmt.Println("Tentativa de recuperar a execução")
	if r := recover(); r != nil {
		fmt.Println(r) //mensagem de panico
		fmt.Println("Execução recuperada com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao() //é chamado antes de matar a aplicação
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	//o valor retornado será FALSE porque é o valor zero do tipo de retorno escolhido para essa função

	panic("[PANIC] A média é exatamente 6!") //deve declarar uma mensagem

}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós Execução")
}
