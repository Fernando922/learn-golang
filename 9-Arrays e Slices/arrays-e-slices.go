package main

import "fmt"



func main(){
	//ARRAY
	var array1[5] string  //posições fixas
	array1[0] = "Posição 1"
	fmt.Println(array1)
	array2 := [5]string{"Posição 1", "2", "3", "4", "5"}
	fmt.Println(array2)


	array3 := [...]int{1,2,3,4,5}  //as posições do array serão definidas pelo quantidade de valores passadas
	fmt.Println(array3)


	//SLICE (não é um array, ele aponta para um array)
	slice := []int{10,11,12,13,14,15,16,17}  //não precisa especificar tamanho do array
	fmt.Println(slice)

	slice = append(slice, 18)  //adiciona no final uma nova posição
	fmt.Println(slice)


	slice2 := array2[1:3]  //pega o primeiro do array e não pega o indice 3
	fmt.Println(slice2)   //[2 3]

	array2[1] = "99"
	fmt.Println(slice2) //[99 3]





	//Arrays Internos
	fmt.Println("---------------")
	slice3 := make([]float32, 10 , 11)  //cria um array de 15 posições e me retorne um slice com os 10 primeiros
	fmt.Println(slice3)
	fmt.Println(len(slice3))  //length   10
	fmt.Println(cap(slice3))  //capacidade  11

	slice3 = append(slice3, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3))  //length   11
	fmt.Println(cap(slice3))  //capacidade  11
	slice3 = append(slice3, 99)
	fmt.Println(slice3)
	fmt.Println(len(slice3))  //length   12
	fmt.Println(cap(slice3))  //capacidade  24  (aumentou a capacidade para receber o novo valor no append, dobrou o valor)


	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))  //length   5
	fmt.Println(cap(slice4))  //capacidade  5
	slice4 = append(slice4, 98)
	fmt.Println(slice4)
	fmt.Println(len(slice4))  //length   6
	fmt.Println(cap(slice4))  //capacidade  12 (dobrou de novo)
	







}