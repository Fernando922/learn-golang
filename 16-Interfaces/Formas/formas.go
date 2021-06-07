package main

import (
	"fmt"
	"math"
)

//Permite a reutilização de funções

type retangulo struct {
	altura float64
	largura float64
}

type circulo struct {
	raio float64
}

type forma interface {
	area() float64
}


func (r retangulo) area() float64 {
	return r.altura * r .largura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

//essa função só pode ser chamadas por structs que possuem o método da interface (forma) no caso o método é area() 
// e não importa qual a sua implementação, mas o nome do método no struct deve ser o mesmo nome do método na interface
// o atributo que você deve declarar na função é o nome da interface desejada

/*
	Resumindo, a função abaixo só recebe structs que atende os requisitos da interface
*/
func escreverArea(f forma){
	fmt.Printf("A area da forma é %0.2f\n", f.area())
}



func main(){
 r := retangulo{10,15}
 escreverArea(r)

 c := circulo{10}
 escreverArea(c)
}