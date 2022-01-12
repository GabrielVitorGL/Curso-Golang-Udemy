package main

import "fmt"

func buscabinaria(n int, items []int) string {
	menorn := 0
	maiorn := len(items) - 1

	for menorn <= maiorn {
		metade := (menorn + maiorn) / 2

		if items[metade] == n {
			return fmt.Sprintf("O número %d foi encontrado na posição [%d]", n, metade)
		} else if items[metade] < n {
			menorn = metade + 1
		} else {
			maiorn = metade - 1
		}
	}
	return "seu número não foi encontrado na lista de elementos"
}

func main() {
	items := []int{1, 3, 4, 27, 34, 39, 64, 75, 132, 135, 185, 723}
	numero := 75

	fmt.Println(buscabinaria(numero, items))
}
