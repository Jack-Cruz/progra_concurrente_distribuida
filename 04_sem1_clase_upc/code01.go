package main

import "fmt"

// Ambito global
const IGV float64 = 18.0

func main() {
	// Ambito local
	// Variable
	var x string = "Primera clase de PCD"
	
	// Conjunto de variables
	var (
		nombre string
		edad int
	)

	nombre = "Juan Carlos"
	edad = 31

	fmt.Println(x)
	fmt.Printf("La edad de %q es %v\n", nombre, edad)
}