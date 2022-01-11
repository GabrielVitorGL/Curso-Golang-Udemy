package main

import "fmt"

// se tiver um slice e quiser passar cada um dos parametros separadamente para uma função variática
// você irá usar 

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"} // slice pois não tem os ... no [] ([...] ou [4] seria um array)
	imprimirAprovados(aprovados...) // colocando os ... voce ira passar cada um dos elementos do slice de forma separada para sua função variática, cada um como um elemento
}
