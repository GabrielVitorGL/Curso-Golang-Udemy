package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// tipos de números sem sinal (positivos): uint8 uint16 uint32 uint64
	var b byte = 255 // byte é um apelido para uint8, nem todos tem apelidos.
	fmt.Println("O byte é", reflect.TypeOf(b))

	// tipos de números com sinal: int8 int16 int32 int64
	i1 := math.MaxInt64 // pegerá o número inteiro de maior valor com 64 bytes
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo do i1 é", reflect.TypeOf(i1))

	var i2 int32 = 'a' // vai mostrar o número da letra 'a' na tabela unicode
	fmt.Println(i2)

	//numeros reais podem ser float32 ou float64
	var x float32 = 49.99 // ou var x = float32(49.99)
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo por padrão de um real é", reflect.TypeOf(49.99)) // o tipo por padrão de um real é float64, se quiser um float 32 tem que determinar


	// boolean (true e false)
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)//contrário de bo, nesse caso o contrario de true vai ser false

	//string
	s1 := "Olá meu nome é Gabriel" // strings sempre dentro de aspas duplas
	fmt.Println(s1 + "!") // Juntara as duas strings
	fmt.Println("O tamanho da string é", len(s1))

	//string com varias linhas
	s2 := `Olá
	meu
	nome
	é
	Gabriel`//CRASE, não é aspas simples
	fmt.Println(s2 + "!")
	fmt.Println("O tamanho da string é", len(s2))

	//char não existe em go, se quiser um unico caractere em uma variavel se usa string
}
