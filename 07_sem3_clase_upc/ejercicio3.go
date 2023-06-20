package main

import "fmt"

// declaración global
var ch chan bool

func procesar(){
	fmt.Println("Mensaje enviado al canal=")
	ch <- true
}

func main(){
	// instancia
	ch = make(chan bool)

	// Canal Sincrono
	go procesar()

	// Leer el valor enviado por el canal
	b := <- ch
	fmt.Println("Mensaje recibido, valor del canal=", b)
}