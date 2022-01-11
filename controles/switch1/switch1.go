package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n) // converteu float para int para ter um valor arredondado da nota
	switch nota {     // na Golang o switch não sai executando todos os cases como no c++ (gloria), então não precisa ficar colocando break (gloria)
	case 10:
		fallthrough // descer e executar o proximo bloco, da mesma forma como o case sem nada no c++
	case 9:
		return "A" // forma 1 de ter um comando para mais de um valor no case
	case 8, 7:
		return "B" // forma 2 de ter um comando para mais de um valor no case
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(11.5))
	fmt.Println(notaParaConceito(-25.61))
}