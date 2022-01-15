package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
	//alterando o nome e o sobrenome pois estamos usando o ponteiro
}

func main() {
	p1 := pessoa{"Pedro", "Silva"}
	fmt.Println(p1.getNomeCompleto())

	//alterando o nome armazenado na struct usando a função que criamos pra isso
	p1.setNomeCompleto("Maria Pereira")
	fmt.Println(p1.getNomeCompleto())
}
