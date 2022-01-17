package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro // isso é um campo anônimo; não vamos colocar o nome identificador, apenas o tipo
	// colocando dessa forma, podemos acessar as variaveis de carro aqui

	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40" // nome é uma variavel de carro, porem estamos acessando pela ferrari pois utilizamos o campo anonimo
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Printf("A ferrari %s está com o turbo ligado? %v\n", f.nome, f.turboLigado)

	fmt.Println(f)
}
