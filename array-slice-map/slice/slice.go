package main

import (
	"fmt"
	"reflect"
)

//? slices servem para "dividir o array em partes" se quiser. Por exemplo, se quiser pegar os elementos de um array apenas do indice 5 em diante, voce usara um slice que vai apontar para o elemento 5 do array e ele conterá os demais elementos de acordo com o número de posições que você quiser
//? slice não é um array
//? O slide define um pedaço de um array
//? Um slice não gera um array diferente, ou seja, ele não aloca novos espaços na memória para os elementos, ele apenas aponta para os elementos que estão armazenados no array
func main() {
	a1 := [3]int{1, 2, 3} //! array
	s1 := []int{1, 2, 3}  //! slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5} // array
	// vamos pegar alguns trechos desse array usando slice
	s2 := a2[1:3] // dessa forma nosso slice conterá os elementos de "a2" do indice 1 ao indice 3, POREM ELE NÃO INCLUE O INDICE 3, desta forma os elementos desse slice seriam: 2, 3
	fmt.Println(a2, s2)

	s3 := a2[:2] // vai do inicio do array até o elemento de indice 2 SEM PEGAR O DE INDICE 2, então fica: 1, 2
	fmt.Println(a2, s3)

	s4 := s2[:1] // podemos fazer um slice de um slice, nesse caso como o s2 era do indice 1 ao 3 SEM INCLUIR O 3 (2, 3), esse sera do inicio ao indice 1 SEM INCLUIR O 1 (2)
	fmt.Println(s2, s4)
}
