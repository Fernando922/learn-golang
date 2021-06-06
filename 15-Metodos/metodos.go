package main

import "fmt"

//Método é uam função que é sempre associada a alguma coisa, interface, struct por exemplo

type usuario struct {
	nome string
	idade uint8
}

//vincula esta função a todo usuario criado
func (u usuario) salvar(){
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool{
	return u.idade >=18
}

//Para alterar os valores do struct instanciado é necessário acessar por ponteiro, já que vamos alterar o struct original e nao uma cópia
func (u *usuario) fazerAniversario() {
	u.idade++  //não precisa desreferenciar

	//é necessário acessar como ponteiro, porque dessa forma você altera o valor do struct original e não da cópia passada como parametro comum
}

func main() {
	usuario1 := usuario{"Fernando", 29}
	usuario1.salvar()


	usuario2 := usuario{"Rogerio", 10}
	usuario2.salvar()
	fmt.Println(usuario2.maiorDeIdade())
	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
