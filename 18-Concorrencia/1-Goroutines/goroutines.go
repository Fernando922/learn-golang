package main

import (
	"fmt"
	"time"
)

func main() {
	// Concorrencia != Parelelismo
	//Paralelismo, tarefas ao mesmo tempo em nucleos diferentes do processador
	//Concorrencia, revezamento de tempo de processamento no mesmo núcleo

	//Go na chamada de um método
	//executa um programa, mas chama o proximo antes dele acabar
	/*
		Você pode ter várias durante a execução do programa
	*/

	go escrever("Olá Mundo") //goroutine (concorrencia)
	escrever("Programando em GO!")

}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
