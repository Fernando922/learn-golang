package main

import "fmt"

func main() {

	//estrutura de chave e valor
usuario := map[string]string{  //dentro dos colchetes é o tipo da chave e fora é o tipo dos valores
	"nome":      "Pedro",
	"sobrenome": "Silva",
}

fmt.Println(usuario["nome"])  //um atributo é acessado dessa forma


usuario2 := map[string]map[string]string {  //um map de chave string onde cada valor dentro dele é um map com chave e valor string
	"nome": {
		"primeiro": "João",
		"ultimo": "Pedro",
	},
	"curso": {
		"nome": "Engenharia",
	},
}

fmt.Println(usuario2)
delete(usuario2, "nome")  //apaga uma chave
fmt.Println(usuario2)
usuario2["signo"] = map[string]string{
	"nome": "Gêmeos",
}
fmt.Println(usuario2)


}

//ANINHAR VARIOS MAPS FICA CONFUSO
