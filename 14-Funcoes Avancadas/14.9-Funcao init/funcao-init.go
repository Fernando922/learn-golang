package main

import "fmt"


//A função init é sempre executada primeiro, antes da função main exclusiva
//Casa arquivo dentro de um pacote pode ter uma função init, diferente da função main que é uma por pacote


//Função init é útil para executar algo antes da main, como configuração de setup, etc.

var n int

func init() {
	fmt.Println("executando a função INIT")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}