package main

import (
	"fmt"
	"time"
)

// nesse programa vamos checar qual o tipo da variavel de um dado passado
func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "tipo não esperado"
	}
}
func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(5))
	fmt.Println(tipo("boa tarde"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
