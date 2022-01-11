package main

import "fmt"

func somar(a int, b int) int { //deve-se passar o tipo de cada variavel que vamos utilizar (as que estao dentro do parenteses) e devemos dizer no final qual o tipo de valor que a função vai retornar, nesse caso int
	return a + b
}

func imprimir(valor int) { //Essa função vai ser igual o que chamamos de "procedimento" no c++, ou seja, uma função que não retorna um valor. Nesse caso também não precisamos dizer o tipo do valor que a função vai retornar, já que ela não vai retornar nada =)
	fmt.Println(valor)
}

func main(){
	resultado := somar(3, 4)
	imprimir(resultado)
}