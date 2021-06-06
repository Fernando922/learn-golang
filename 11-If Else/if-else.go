package main

import "fmt"

func main() {
	numero := 10

	minhaidade := 29
	idadeDela := 28

	if numero > 15 { //nao precisa colocar parentes, é obrigatório colocar chaves
		fmt.Println("Maior que 15")
	} else if numero == 0 {
		fmt.Println("Número é igual a zero")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	//IF INIT

	//declarou a variável e usou na mesma linha
	//o escopo dela é dentro do if, fora do if ela não existirá
	//se for usar uma variável apenas pra isso, é interessante fazer o if init
	//tipo de quiser comparar se a soma de outras variáveis representa algo
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	}

	if somaDasIdades := minhaidade + idadeDela; somaDasIdades > 50 {
		fmt.Println("É maior que 50", somaDasIdades)
	}

	//A variável morre no escopo do If
	//fmt.Println(somaDasIdades)    //vai dar erro

}
