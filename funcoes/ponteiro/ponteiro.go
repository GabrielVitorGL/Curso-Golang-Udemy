package main

import "fmt"

/* um ponteiro em uma função é muito útil pois quando passamos uma variável para uma função,
estamos passando uma cópia dela e não mexendo no valor original.
se quisermos , dentro de uma função, alterar o valor original de uma variável fora dela,
iremos utilizar o ponteiro.*/

func inc1(n int) {
	n++
}

func inc2(n *int) {
	// o * no int é para ter acesso ao valor que o ponteiro aponta e não do próprio endereço da memória

	*n++
}

func main() {
	n := 1

	inc1(n) // vai operar com a cópia desse valor
	fmt.Println(n)

	// o & é usado para passar o endereço de uma variável
	inc2(&n) // vai operar com o valor original
	fmt.Println(n)
}
