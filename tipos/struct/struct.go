package main

import "fmt"

//struct é um bloco que agrupa várias váriaveis, usaremos a seguinte sintaxe:

type produto struct { // assim se declara um tipo (nesse caso do tipo struct)
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver

func (p produto) /*dentro da função vamos chamar "produto" por "p"*/ precoComDesconto() float64 {
	//nome da função (obs: não precisamos passar nada como parâmetro pois todas as informações que precisamos para calcular o preco com desconto já estão dentro do tipo "produto", que é o nome, preco e desconto)
	//p.preco assim vamos acessar a variavel "preco" da struct "produto" que referenciamos por "p"
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto // já que criamos o tipo produto, vamos utilizar ele aqui
	produto1 = produto{
		// aqui sim iremos utilizar ,
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"Notebook", 1989.90, 0.10} // criando a struct de forma reduzida
	fmt.Println(produto2, produto2.precoComDesconto())
}
