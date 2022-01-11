package main

import "fmt"

// funções variáticas são funções que podem receber um número variável de parametros

func media(numeros ...float64) float64 { // os ... significam que isso é uma função variática, então podemos passar quantos parametros precisarmos
	// uma função variática, trabalha com os parametros em uma array, entao precisamos percorrer essa array para pegar os valores
	total := 0.0
	for _, num := range numeros { // vamos percorrer o array numeros para ir somando os elementos dele na variável total; usamos o _ para ignorar o indice já que não precisamos dele
		total += num
	}
	return total / float64(len(numeros)) // total das notas dividido pelo tamanho do array
}

func main() {
	fmt.Printf("Média: %.2f", media(7.7, 8.1, 5.9, 6.4, 9.9))
}
