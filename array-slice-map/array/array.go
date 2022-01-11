package main

import "fmt"

func main() {
	//se você criar um array do tipo int, só poderá ter valores inteiros nele;
	//não é possivel mudar o tamanho de um array criado com x posições
	// o indice começa no 0 assim como no c++
	var notas [3]float64 //assim se cria um array; nome[posicoes]tipo
	fmt.Println(notas)   //os valores padrão de um array vazio é "0 0 0" nesse caso que definimos como 3 posiçoes

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1 // voce pode definir os valores todos em uma linha se quiser
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas)) // para calcular a media, precisamos do tamanho do array, porem as notas estão em float e o tamanho do array em int. Então precisamos converter o tamanho do array para float para podermos fazer a conta.
	fmt.Printf("Média %.2f\n", media)
}
