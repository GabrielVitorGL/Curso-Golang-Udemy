package main

import "fmt"

// se chamarmos uma função em que nela x era igual a 10, porém em outro escopo em que x é 20
// ela irá ter o valor de x como 10, pois foi onde ela foi criada


func closure() func() { // uma função que irá retornar outra função
	x := 10
	var funcao = func() { // criando a função que vamos retornar
		// aqui é o conteúdo da função "funcao" que está dentro da função "closure"
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)

	imprimeX := closure() // como a função "closure" retorna uma função, "imprimeX" será declarado como uma função
	imprimeX()
}
