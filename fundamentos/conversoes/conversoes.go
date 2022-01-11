package main

import (
	"fmt"
	"strconv"
)

func main() {
	//vamos aprender a converter os tipos pois mesmo dois valores numericos de tipos diferentes não funcionam, precisamos converter os dois para o mesmo tipo antes

	x := 2.4
	y := 2
	//fmt.Println(x / y) não funciona pois as variaveis sao de tipos diferentes
	fmt.Println(x / float64(y)) //fazendo dessa forma convertemos para funcionar

	nota := 6.9
	notaFinal := int(nota) // Convertendo para int a nota
	fmt.Println(notaFinal)

	// Como converter inteiro para string:
	fmt.Println("Teste" + strconv.Itoa(123))

	//String para int
	num, _ := strconv.Atoi(("123")) // retorna dois valores dessa conversao, o primeiro que vai ser armazenado em "num" será o número, e o segundo "_" será um erro caso você pase um valor errado (o "_" deleta o valor armazenado nele)
	fmt.Println(num - 122)

	// o pacote strconv tambem converte o tipo boolean, veja so
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}
