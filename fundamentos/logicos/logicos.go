package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) { // voce pode dizer so uma vez o tipo das variaveis caso queria que as duas sejam do mesmo tipo; os tres tipos no final dizem quais os tipos dos valores que vamos retornar dessa função, no caso como vamos retornar 3 valores, percisamos colocar 3 vezes
	comprarTv50 := trab1 && trab2 // se quiser que sejam true para acontecer, não precisa dizer trab1 == true por exemplo, só colocar elas la mesmo
	comprarTv32 := trab1 != trab2 // nesse caso precisariamos usar um "OU exclusivo (só pode acontecer uma das duas(diferente do caso de ComprarSorvete))", como não temos na linguagem Golang vamos usar o != que vai ter a mesma função (se trab1 for false e trab2 for true ou vice-versa, será satisfeito o "ou")
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true) // no true true é onde vamos passar o valor das duas variaveis que usamos na função. Como esse programa ainda não pede inputs do usuario, estamos colocando manualmente aqui. Basta mudar o true true para testar as opções que temos
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Sem sorvete: %t", tv50, tv32, sorvete, !sorvete) // para descobrir se "sem sorvete" vai ser true ou false, negamos a variavel "sorvete". Ou seja, se sorvete for true, "sem sorvete" será falso e vice-versa
}
