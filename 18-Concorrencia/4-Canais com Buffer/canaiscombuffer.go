package main

import "fmt"

func main() {
	//canal com buffer só bloqueia quando atinge sua capacidade maxima
	canal := make(chan string, 2) //canal com buffer

	canal <- "Olá Mundo!"
	canal <- "Programando em GO!"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
