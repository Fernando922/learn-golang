package formas

import (
	"math"
)

//Permite a reutilização de funções

type Retangulo struct {
	altura  float64
	largura float64
}

type Circulo struct {
	raio float64
}

type Forma interface {
	Area() float64
}

func (r Retangulo) Area() float64 {
	return r.altura * r.largura
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}
