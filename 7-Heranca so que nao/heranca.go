package main

import "fmt"


type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}


type estudante struct {
	pessoa  //herda tudo de pessoa
	curso string
	faculdade string
}



func main(){
	fmt.Println("Herança")


	p1 := pessoa{"João", "Pedro", 20, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1.nome)


	e2 := estudante{pessoa{"Fernando", "de Paula", 29,175}, "Analise de Sistemas", "Fatec" }
	fmt.Println(e2.sobrenome)
}