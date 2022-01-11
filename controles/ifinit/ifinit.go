package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	return s.Intn(10)
}

func main() { // vamos ja inicializar a variavel i e usa-la para uma comparação no nosso if
	if i := numeroAleatorio(); i > 5 { // a variavel i vai ser temporaria e só existira dentro do bloco do if
		fmt.Println("Ganhou!")
	} else {
		fmt.Println("Perdeu!")
	}
}
