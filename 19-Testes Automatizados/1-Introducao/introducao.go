package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	TipoDeEndereco := enderecos.TipoDeEndereco("Rodovia Paulista")
	fmt.Println(TipoDeEndereco)
}
