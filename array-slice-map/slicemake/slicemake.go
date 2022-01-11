package main

import "fmt"

func main() {

	//? Quando se cria um slice sem referenciar um array (com o metodo make), internamente ele criará um array com o número de posições do slice

	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)        //? nesse caso, o slice terá 10 posições porém o array interno que será criado terá 20 posições
	fmt.Println(s, len(s), cap(s)) // len = tamanho do slice, cap = capacidade máxima de elementos

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) // vamos preencher as casas restantes do nosso slice com a função append, que preenche a partir das casas vazias
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
	//? Quando se atinge a capacidade máxima do array interno, a capacidade máxima vai dobrar de tamanho para você continuar trabalhando
}
