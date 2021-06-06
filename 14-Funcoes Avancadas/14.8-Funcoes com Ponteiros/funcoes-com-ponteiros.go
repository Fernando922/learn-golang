package main

import "fmt"


// * acessa o valor que o ponteiro aponta
// & acessa a referencia de memória



//Se você quer que uma variável seja alterada depois de chamar uma função
// é necessário fazer a alteração como ponteiro e não como cópia de um valor


//Não muda o valor da variável original (é passado a cópia de um valor)
func inverterSinal(numero int) int {
	return numero * -1
}

//muda a variável inicial (é passado a referencia da variável original)
func inverterSinalComPonteiro(numero *int){
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)

}
