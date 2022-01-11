package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{ // forma de inicializar um map já com os elementos
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0, // O último elemento também precisa ter vírgula
	}

	funcsESalarios["Rafael Filho"] = 1350.0
	delete(funcsESalarios, "Inexistente") // nao tem problema deletar um elemento que nao existe

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
