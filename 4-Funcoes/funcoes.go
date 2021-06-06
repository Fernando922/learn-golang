package main

import "fmt"

//Função simples, os atributos que ela recebem devem ser tipados, e o retorno também
func somar(n1 int8, n2 int8) int8 {
	return n1+n2
}

//se os atributos são do mesmo tipo, pode-se tipar no final, ou seja, quem não foi tipado no início pega a tipagem do ultimo atributo passado
// essa função tem multiplos retornos, vantagem do go, a ordem dos atributos deve ser a mesma da ordem dos retornos
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 -n2
	return soma, subtracao	//dois retornos por função
}


func main(){
	soma := somar(10,20)
	fmt.Println(soma)  //30


	//atribuindo uma função anonima para uma variável, deve-se declarar como func() "não passar o nome em seguida"
	var f = func() {
		fmt.Println("Função F")
	}

	f()  //Função F


	//Função anonima que recebe uma string e retorna uma string
	var f2 = func(txt string) string{
		fmt.Println(txt)
		return txt
	}

	f2("Fernando") //Fernando (por causa do println)

	resultado := f2("Texto da função 2")  //Texto da função 2
	fmt.Println(resultado)  //Texto da função 2


	//Ignorando valores retornados
	//Como no GO toda variável declarada deve ser utilizada, se você não quiser usar um dos retornos específicos, deve-se declarar com "underline"
	resultadoSoma, _ := calculosMatematicos(10,15)  //underline significa que o retorno será ignorado
	fmt.Println(resultadoSoma)
}