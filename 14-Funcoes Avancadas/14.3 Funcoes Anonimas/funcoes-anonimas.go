package main

import "fmt"

//a função é criada e já executada em seguida
func main() {
	retorno := func(texto string)string {
		fmt.Println("Olá Mundo", texto)
		return fmt.Sprintf("Recebido: %s", texto)  //formata strings
	}("Fernando")
	fmt.Println(retorno)
}
