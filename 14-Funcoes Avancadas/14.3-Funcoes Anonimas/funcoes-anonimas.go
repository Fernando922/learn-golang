package main

import "fmt"


//Uma função anonima é atribuida a uma variávell e poder ser executada automaticamente
//a função é criada e já executada em seguida
func main() {
	retorno := func(texto string)string {
		fmt.Println("Olá Mundo", texto)
		return fmt.Sprintf("Recebido: %s", texto)  //formata strings
	}("Fernando")
	fmt.Println(retorno)
}
