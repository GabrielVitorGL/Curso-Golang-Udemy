package main

import "runtime/debug"

//Exemplo de funções dentro de outras funções

func f1() {
	f2()
}

func f2() {
	f3()
}

func f3() {
	debug.PrintStack()
}

func main() {
	f1()
}
