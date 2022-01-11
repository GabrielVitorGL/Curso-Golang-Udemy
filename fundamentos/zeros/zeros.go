package main

import "fmt"

func main() {
	//vamos descobrir qual o valor inicial de cada tipo de variavel
	//ja que não vamos definir nenhum valor inicial, precisamos dizer o tipo
	var a int
	var b float64
	var c bool
	var d string
	var e *int

	fmt.Printf("%v %v %v %q %v",a,b,c,d,e)
}