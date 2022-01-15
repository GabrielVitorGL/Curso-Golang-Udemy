package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item //pegamos elementos da struct "item", estamos armazenando um "tipo dentro de outro"
}

func (p pedido) valorTotal() float64 { // não precisa de parametro pois já temos tudo nos tipos
	total := 0.0

	for _, item := range p.itens { // ignorando o indice
		total += item.preco * float64(item.qtde) // temos que converter a quantidade para fazer a conta já que ela é do tipo int
	}
	return total
}

func main() {
	pedido := pedido{ // prefira escrever mais e deixar mais claro do que escrever menos e deixar o código confuso;
		//essa notação é maior mas deixa o código mais claro
		userID: 1,
		itens: []item{
			item{1, 2, 12.10}, //forma reduzida porém nem tanto recomendado pois fica mais confusa
			item{2, 1, 23.49},
			item{11, 100, 3.13}, //declaramos 3 itens como exemplo
		}, //precisamos colocar a , aqui para fechar a parte dos "itens"
	}

	fmt.Printf("Valor total do pedido é: R$%.2f", pedido.valorTotal())
}
