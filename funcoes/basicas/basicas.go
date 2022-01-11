package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string { // forma otimizada de criar a função f2
	return fmt.Sprintf("F4: %s %s", p1, p2) // Sprintf não imprime no console, e sim retorna uma string formatada.
}

func f5() (string, string) { // como retornar 2 valores de uma função
	return "Retorno1", "Retorno 2"
}

func main() {
	f1()
	f2("Param1", "Param2")
	r3, r4 := f3(), f4("Param1", "Param2")
	fmt.Println(r3)
	fmt.Println(r4)

	r51, r52 := f5()
	//_, r52 := f5()
	//se quiser ignorar um valor por exemplo
	fmt.Println("F5:", r51, r52)

}
