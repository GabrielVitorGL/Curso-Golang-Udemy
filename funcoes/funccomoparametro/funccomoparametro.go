package main

import "fmt"

//como passar uma função como parametro para outra função
//um exemplo de uso seria criar uma função que executará um bloco com várias outras funções, sendo só preciso chamá-la uma vez

func multiplicacao(a, b int) int {
	return a * b
}

func exec(funcao func(int, int) int, p1, p2 int) int { // funcao func(int, int) é o nome da função, o tipo (func) e os tipos de dados da função que passamos, no caso a e b eram int; depois você cria dois parametros para armazenar os valores, no caso p1 e p2 e da o tipo deles, e por último o tipo de valor que você quer retornar
	return funcao(p1, p2)
}

func main() {
	resultado := exec(multiplicacao, 3, 4)
	fmt.Println(resultado)
}
