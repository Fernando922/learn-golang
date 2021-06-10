package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Sincronizar goroutines
*/

func main() {

	//criar grupo de espera
	var waitGroup sync.WaitGroup

	//quantidade de goroutines ao mesmo tempo
	waitGroup.Add(2)

	go func() {
		escrever("Olá Mundo")
		waitGroup.Done() //tira -1 do contador do waitgroup
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() //tira -1 do contador do waitgroup
	}()

	waitGroup.Wait() //espera a contagem de goroutines ficar em zero para continuar a execução
	fmt.Println("Fim da execução")

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
