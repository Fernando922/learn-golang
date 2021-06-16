// Teste de unidade
//testa uma pequena parte do código

package enderecos_test

import (
	. "introducao-testes/enderecos"  //alias que representa tudo de enderecos
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

//t é um ponteiro de testing.T
//função começa com "Testxxxx"
//é boa pratica que o nome se o mesmo da função testada
func TestTipoDeEndereco(t *testing.T) {

	t.Parallel()  //permite que rode em paralelo

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida ABC", "Avenida"},
		{"Rodovia ABC", "Rodovia"},
		{"RodoVia ABC", "Rodovia"},
		{"Estrada ABC", "Estrada"},
		{"Est ABC", "Tipo Inválido!"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
				cenario.retornoEsperado,
				retornoRecebido)
		}
	}

}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Error("Teste quebrou")
	}
}
