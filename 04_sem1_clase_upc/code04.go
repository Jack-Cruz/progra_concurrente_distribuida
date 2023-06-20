package main

import "fmt"

func main() {
	var num float64
	fmt.Print("Ingresa feet: ")
	fmt.Scanf("%f", &num)

	fmt.Printf("En metros es %v metros\n", num*0.3048)

}