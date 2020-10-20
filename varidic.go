package main

import "fmt"

func mayor(a... int) int { 
	mayor := 0
	for i := 0; i < len(a); i++{
		if a[i] > mayor {
			mayor = a[i]
		}
	}
	return mayor
}

func main(){
	a := []int{1,4,2,8,0,10}
	
	fmt.Println(mayor(a...)) // cada elemento del slice se envía como parámetros
}