package main

import "fmt"

func fatorial(n int) (int, error) {
	if n < 0 {
		return -1, fmt.Errorf("número inválido: %d", n)
	}
	resultado := 1
	for n = n; n > 0; n-- {
		resultado *= n
	}
	return resultado, nil
}

func main() {
	resultado, erro := fatorial(5)
	if erro != nil {
		fmt.Println(erro)
	} else {
		fmt.Println(resultado)
	}
}
