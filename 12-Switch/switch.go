package main

import "fmt"


func diaDaSemana1(numero int)string{

	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5: 
		return "Quinta"
	case 6: 
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Dia não encontrado"
	}

}

//NAO PRECISA USAR O BREAK

func diaDaSemana2(numero int)string{
	var diaDaSemana string
	switch {
	case numero == 1 :
		diaDaSemana= "Domingo"
		fallthrough  //faz cair na proxima condição e retorna o resultado
	case numero == 2 :	
		diaDaSemana= "Segunda"
	default:
		diaDaSemana="Não encontrado"
	}
	return diaDaSemana
}

func main(){
	fmt.Println(diaDaSemana1(3))  //terça
	fmt.Println(diaDaSemana2(1))  //segunda
}