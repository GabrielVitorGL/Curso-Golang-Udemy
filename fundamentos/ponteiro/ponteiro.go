package main

import "fmt"

func main() {
	i := 1

	var p *int = nil // forma de criar um ponteiro

	p = &i // colocando o endereço da variavel "i" na variavel "p"
	*p++   // colocando o asterisco antes do p, vamos estar mudando o valor da variavel no qual ele esta apontando. Nesse caso, vamos estar somando 1 na variavel "i" utilizando o ponteiro dela.
	i++

	fmt.Println(p, *p, i, &i)

	//p++ por exemplo não é permitido, já que p está armazenando o endereço da variavel "i", e não podemos fazer operações com esses endereços
}
