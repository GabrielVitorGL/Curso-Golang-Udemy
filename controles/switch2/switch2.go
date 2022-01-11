package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { // dessa forma (sem passar uma variavel no switch) ele vai procurar qual case retorna verdadeiro, o primeiro que retornar ele executará e sairá do switch.
	//switch true{} também teria o mesmo resultado
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}
}
