package main

import (
	"fmt"
	"math"
)

func main() { // TODOS OS OPERADORES COMUNS SÃO IGUAIS DO C++
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Módulo =>", a%b)

	//Outros operadores
	//bitwise (pega o valor binario das variaveis passadas, e compara bit a bit (ex: 11 & 10 = 10 >> 1 & 1 = 1, 1 & 0 = 0(parecido com regra de sinal)))
	fmt.Println("E =>", a&b)   //11 & 10 = 10
	fmt.Println("Ou =>", a|b)  // 11 | 10 = 11
	fmt.Println("Xor =>", a^b) // 11 ^ 10 = 01 //!OBS: isso não é potência

	//outras operações serão possiveis usando o math

	c := 3.0
	d := 2.0 // declaramos dessa forma pois já sao definidas como float, já que a maioria dos comandos do math precisam que os valores sejam desse tipo

	fmt.Println("Maior =>", math.Max(float64(a), float64(b))) // exemplo com as variaveis do tipo int
	fmt.Println("Menor =>", math.Min(c, d))                   // com as que já são do tipo float não precisamos converter
	fmt.Println("Exponenciação =>", math.Pow(c, d))
	fmt.Println("Raiz quadrada =>", math.Sqrt(c))

}
