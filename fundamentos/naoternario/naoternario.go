package main

import "fmt"

//no go nao temos operadores ternarios, mas temos outras formas de fazer isso

func obterResultado(nota float64) string {// primeiro dizemos a variavel que vamos trazer e seu tipo, e fora do parenteses dizemos o tipo de valor que essa função retornará; Se quisermos que a função retorne mais de um tipo, colocamos o tipo de cada um dos valores retornados dentro de parenteses, e separados por virgula.
	if nota >= 6{
		return "Aprovado"
	}
	return "Reprovado" // nao precisamos colocar esse return no else, já que se a nota for >= 6 vai entrar no if e assim que ele executar o return, ele já sai da função. Se quiser retornar mais de um valor, usam-se todos no mesmo return
	// Se a Golang suportasse operadores ternarios, seria assim a forma como nós fariamos os codigos acima:
	//! return nota>=6 ? "Aprovado" : "Reprovado"
}

func main() {
	fmt.Println(obterResultado(6.2))
}