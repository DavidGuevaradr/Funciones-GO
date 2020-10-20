package main

import "fmt"

func generadorimPares() func() int {
	i := int(1) // i permanecerá en el clousure de la función anónima a retornar
	return func() int {
		var impar = i
		i += 2
		return impar
	}
}

func main() {
	nextimPar := generadorimPares()
	fmt.Println(nextimPar())
	fmt.Println(nextimPar())
	fmt.Println(nextimPar())
	fmt.Println(nextimPar())
	fmt.Println(nextimPar())
}