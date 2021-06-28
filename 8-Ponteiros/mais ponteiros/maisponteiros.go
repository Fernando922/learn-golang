package main

import "fmt"

func main() {
	var creature string = "shark"
	var pointer *string = &creature

	fmt.Println("creature =", creature) // creature = shark
	fmt.Println("pointer =", pointer)   // pointer = 0xc000044240

	fmt.Println("*pointer =", *pointer) //  *pointer = shark (com o asterisco voce le o valor real)

	*pointer = "jellyfish"
	fmt.Println("*pointer =", *pointer) //*pointer = jellyfish (dessa forma eu consigo alterar o valor original)
}

/*

Decidir entre o momento de enviar um ponteiro - ao invés de enviar um valor - é predominantemente uma questão de saber se você deseja ou não que o valor mude.
Se não quiser que o valor mude, envie-o como um valor. Se quiser que a função para a qual você está enviando sua variável possa alterá-la, então você a enviaria como um ponteiro.
*/
