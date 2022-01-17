package main

import "fmt"

type imprimivel interface { // tipo interface
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{
		nome:  "Calça Jeans",
		preco: 79.90,
	}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	p2 := produto{"Camiseta", 50.00}
	imprimir(p2)
}
