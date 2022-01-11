package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	//array1 = append(array1, 4, 5, 6)
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1) // copiando o slice1 para o slice2, como o slice2 só tem 2 posições, pegou os 2 primeiros valores do slice1
	//? O copy só funciona com slice's, com array's não funcionam
	//? O copy não expande o tamanho do slice se necessário
	fmt.Println(slice2)
}
