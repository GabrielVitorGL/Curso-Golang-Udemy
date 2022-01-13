package main

import "fmt"

// a função "init" será sempre executado por padrão antes da função main, mesmo sem chamá-la

func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}
