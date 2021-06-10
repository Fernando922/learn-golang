package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá Mundo!", canal)

	// for {
	// 	mensagem, aberto := <-canal //recebe um valor do canal, só continua a execução quando receber um valor
	// 	//se um canal estiver esperando receber e nunca mais receber nada é estourado um deadlock
	// 	//deadlock não é pego em compilação, só em execução, cuidado pra nao enviar para prod

	// 	//verifica se o canal está fechado
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	//Maneira simplificada, enquanto estiver aberto o canal vai chamar essa função
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto //atribui para o canal a mensagem
		time.Sleep(time.Second)
	}

	close(canal) //fecha o canal após o termino do loop
}
