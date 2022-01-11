package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{ // um map cujo valor é outro map
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   8456.78,
		},
		"J": {
			"José João": 11325.45,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}
	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

	fmt.Printf("\n\n")

	for _, map2 := range funcsPorLetra {
		for nome, salario := range map2 {
			fmt.Println(nome, salario)
		}
	}
}
