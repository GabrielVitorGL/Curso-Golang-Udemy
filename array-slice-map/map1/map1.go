package main

import "fmt"

func main() {
	// var aprovados map[int]string
	//maps tem que ser inicializados
	aprovados := make(map[int]string) // nesse caso a chave será de valor inteiro, e conteúdo sera do tipo string

	aprovados[12345678978] = "Maria"
	aprovados[98765432100] = "Pedro"
	aprovados[95135745682] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados { // For para percorrer um map
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[95135745682]) // A chave é como a identificação do map, e dessa forma vamos obter o conteudo, no caso "ana"
	delete(aprovados, 95135745682)      // como deletar alguma chave do map
	fmt.Println(aprovados[95135745682])
}
