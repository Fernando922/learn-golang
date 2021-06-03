package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	auxiliar.Escrever()
	fmt.Println("Escrevendo do arquivo main")
	erro := checkmail.ValidateFormat("dipaula018@gmail.com") //atribui retorno a uma variável
	fmt.Println(erro)

}
