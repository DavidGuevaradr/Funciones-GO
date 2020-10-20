package main

import "fmt"


func main() {
	change := func(a *int, b *int) *int{

		aux := *a
		*a = *b
		*b = aux
		fmt.Println("\n\nintercambio Realizado")

		return a
		return b
	
	}

	a := 10
	b := 20

	fmt.Print("\nletra a con valor: ", a)
	fmt.Print("\nletra b con valor: ", b)
	change(&a, &b)
	fmt.Print("\nletra a con valor: ", a)
	fmt.Print("\nletra b con valor: ", b)

}
