package main

import "fmt"

//provando que realmente dois slices de uma mesma array realmente apontam para a mesma array

func main() {

	//! MINHA SOLUÇÃO
	a1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s1 := a1[:6]
	s2 := a1[5:]
	fmt.Println(a1, s1, s2)

	var pa1, ps1, ps2, pa2 *int = &a1[5], &s1[5], &s2[0], &a2[5]
	fmt.Println(pa1, ps1, ps2, pa2)

	//! SOLUÇÃO DA AULA
	ss1 := make([]int, 10, 20)
	ss2 := append(ss1, 1, 2, 3)
	fmt.Println("")
	fmt.Println(ss1, ss2)

	ss1[0] = 7
	fmt.Println(ss1, ss2)
}
