package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"` //qual sera o nome da chave após a conversão
	Raca  string `json:"raca"`
	Idade uint   `json:"-"`  //ignora este valor no retorno
}

func main() {
	cachorroEmJSON := `{"nome":"Rex","raca":"Dálmata","idade":3}`

	var c cachorro //será alterada tb logo após a proxima linha

	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil { //endereço de memoria para onde vao os dados
		log.Fatal(erro)
	}

	fmt.Println(c)

	cachorro2EmJSON := `{"nome":"Rex","raca":"Dálmata"}`

	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil { //endereço de memoria para onde vao os dados
		log.Fatal(erro)
	}

	fmt.Println(c2)

}
