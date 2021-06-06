package main

import "fmt"

//Defer = adiar

/*
	Adiar a execução dessa função até o último momento possível caso seja função sem retorno
	Executa antes do return caso a função tenha return

	É util quando há consulta em banco de dados

	exemplo fechar a conexão com o banco, dando certo ou não

*/

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada, resultado será retornado")  //neste caso será retornado antes do return, evitando duplicar código
	fmt.Println("Entrando na função se o aluno está aprovado")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	defer funcao1() //será executado no final de tudo
	funcao2()
	fmt.Println(alunoEstaAprovado(7, 8))
}
