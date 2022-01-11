package main

import (
	"fmt"
	"time"
)

func main() {
	// atribuição: 1 igual
	// comparação: 2 iguais

	// todos os exemplos a seguir vao retornar true ou false
	fmt.Println("Strings:", "Banana" == "Banana") // Você pode compara strings
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas:", d1.Equal(d2)) // outra forma de checar se o valor de duas variaveis sao iguais

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{Nome: "João"}
	p2 := Pessoa{Nome: "João"}
	fmt.Println("Pessoas:", p1 == p2)
}
