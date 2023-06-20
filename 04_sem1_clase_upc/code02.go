package main

import "fmt"

func main() {
	var num float64
	fmt.Print("Ingrese un número: ")
	fmt.Scanf("%f", &num)

	fmt.Printf("El doble de %v es %v\n", num, 2*num)
}