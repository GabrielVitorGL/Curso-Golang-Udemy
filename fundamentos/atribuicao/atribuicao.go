package main

import "fmt"

func main() {
	var b byte = 3 // forma 1 de atribuir valor
	fmt.Println(b)

	i := 3 // forma 2 de atribuir valor (bem mais facil)
	fmt.Println(i)
	i += 3 // i = i + 3
	fmt.Println(i)
	i -= 3 // i = i - 3
	fmt.Println(i)
	i /= 2 // i = i / 2
	fmt.Println(i)
	i *= 2 // i = i * 2
	fmt.Println(i)
	i %= 2 // i = i % 2
	fmt.Println(i)

	//da pra atribuir valor a varias variaveis de uma vez só, dessa forma:
	x, y := 1, 2 // 1 pra x; 2 pra y
	fmt.Println(x, y)

	//da pra substituir o valor de duas variáveis de forma bem fácil, desse jeito:
	x, y = y, x
	fmt.Println(x, y)
}
