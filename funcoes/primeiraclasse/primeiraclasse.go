package main

import "fmt"

//função dentro de uma variável

var soma = func(a, b int) int {
	return a + b
}

func main() {
	//chamamos uma função dentro de uma variável da mesma forma como uma função nomeada
	fmt.Println(soma(2, 3))

	//você pode criar uma função dentro do main dessa forma
	sub := func(a, b int) int {
		return a - b
	}
	fmt.Println(sub(2, 3))
}
