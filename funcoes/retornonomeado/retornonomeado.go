package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return // se declarou la em cima já com os nomes, vai retornar na ordem que está la
	//esse retorno chama retorno limpo
}

func main() {
	x, y := trocar(4, 13)
	fmt.Println(x, y)
}
