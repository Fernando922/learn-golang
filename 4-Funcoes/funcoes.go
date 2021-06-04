package main

import "fmt"


func somar(n1 int8, n2 int8) int8 {
	return n1+n2
}


func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 -n2
	return soma, subtracao	//dois retornos por função
}


func main(){
	soma := somar(10,20)
	fmt.Println(soma)


	var f = func() {
		fmt.Println("Função F")
	}

	f()

	var f2 = func(txt string) string{
		fmt.Println(txt)
		return txt
	}

	f2("Fernando")

	resultado := f2("Texto da função 2")
	fmt.Println(resultado)


	resultadoSoma, _ := calculosMatematicos(10,15)  //underline significa que o retorno será ignorado
	fmt.Println(resultadoSoma)
}