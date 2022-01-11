package main

import (
	"fmt"
	"math" // Se quiser usar "m" para chamar math, por exemplo, ficaria:
	//m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // se já receber valor não precisa dizer o tipo
	/*
		const(
			a = 1
			b = 2
		)
		var(
			c = 3
			d = 4
		)
		*formas de declarar mais de uma variavel:
		?var e, f bool = true, false
		ou
		!g, h, i := 2, false, "epa!"
		!mais recomendado
	*/

	//forma reduzida de criar uma variavel e já atribuir um valor a ela:
	area := PI * math.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)
}
