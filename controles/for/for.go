package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1

	// nao se usa parenteses no for
	for i <= 10 { // assim seria exatamente igual ao while
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 { // for da maneira convencional, inicializamos a variavel, passamos uma condição e um incremento.
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando... ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	//laço infinito:
	for {
		fmt.Println("Laço infinito")
		time.Sleep(time.Second)
	}
}
