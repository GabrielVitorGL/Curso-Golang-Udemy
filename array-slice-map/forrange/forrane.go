package main

import "fmt"

//como percorrer os arrays de uma forma mais otimizada, com esse comando vamos percorrer todos os elementos e podemos salvar os indices e os elementos

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // dessa forma voce inicializa um array da forma simplificada, e os ... querem dizer que o compilador vai determinar o numero de indices desse array de acordo com o numero de valores que passamos pra ele

	for i, numero := range numeros { // usando esse "range numeros", vai retornar dois valores (armazenados em i e numero); O primeiro valor será o indice e o outro o elemento desse indice
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros { // caso voce nao queira armazenar os indices
		fmt.Println(num)
	}
}
