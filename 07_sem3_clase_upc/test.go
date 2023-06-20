package main

import "fmt"

var turn int = 1

func p(ch chan int) {
	for {
		// Código sección no crítico (SNC)
		// Código crítico (SC)
	}
}

func q(ch chan int) {
	for {
		// Código SNC
		// Código SC
	}
}

func main() {
	ch = make(chan int)
	go p(ch)
	go q(ch)
}