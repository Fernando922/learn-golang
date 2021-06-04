package main

import "fmt"

//Coleção de campos, cada nome tem um campo e um tipo
//É o mais perto de objeto em go



//Um struct pode ser criado e não utilizado

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint8
}

func main(){
	fmt.Println("Arquivo structs")


	var u usuario
	fmt.Println(u)  //exibe os valores "ZERO" string vazia e numero zero
	u.nome = "Fernando"
	u.idade = 29
	fmt.Println(u)
	enderecoExemplo := endereco{"Rua Qualquer", 0}
	usuario2 := usuario{"Josias", 25, enderecoExemplo}    //ABRE CHAVES!!!
	fmt.Println(usuario2)



	usuario3 := usuario{nome: "Rogério"}  //quando você especifica o nome do campo você não precisa declarar todos
	fmt.Println(usuario3)
}