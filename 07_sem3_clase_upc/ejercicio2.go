package main

import "fmt"

// declaración global
var ch chan bool

func main(){
	// instancia
	ch = make(chan bool)

	// Canal Sincrono
	ch <- true

	// Leer el valor enviado por el canal
	b := <- ch
	fmt.Println("Mensaje recibido, valor del canal=", b)
}