package main

import "fmt"

// esse programa mostrará o fatorial de um número mas não é a forma mais eficiente de fazer isso
// desafio: utilizar um recuso que já aprendemos para descobrir o fatorial de um número de uma forma melhor
// farei na pasta desafiofatorial

func fatorial(n int /*recebe um numero inteiro*/) (int, error /*retorna um int e um erro*/) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("número inválido: %d", n) // retornamos -1 para dizer que é um valor inválido e uma mensagem de erro
	case n == 0:
		return 1, nil // se o numero for 0, o fatorial será 1 e "nil" é porque não vamos retornar um erro
	default:
		fatorialAnterior, _ := fatorial(n - 1) // chamamos a própria função que estamos, e vamos cada vez multiplicando pelo fatorial anterior já que passamos n - 1
		return n * fatorialAnterior, nil
		// então, se passarmos 8, a conta vai ficar 8*7*6*5*4*3*2*1
		// ele vai ficar entrando aqui diversas vezes até o n ser 0
		/* vai ser algo do tipo:
		8
			7
				6
					5
						4
							3
								2
									1
										1
										sai
									sai
								sai
							sai
						sai
					sai
				sai
			sai
		termina a função*/
	}
}

func main() {
	resultado, _ := fatorial(5) // passando 5 para nossa função
	fmt.Println(resultado)
	_, erro := fatorial(-2) // passando -2 para ver o erro
	if erro != nil {
		fmt.Println(erro)
	}
	/* a forma mais comum que utilizamos para trabalhar com erro seria assim:

	resultado, erro := fatorial (x)
	if erro != nil {
		fmt.Println(erro)
	}
	else{
		fmt.Println(resultado)
	}

	*/
}
