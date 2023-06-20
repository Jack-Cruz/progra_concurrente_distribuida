package main

import (
	"fmt"
	"time"
)

var turn int = 1

func p() {
	for {
		// Código sección no crítico (SNC)
		fmt.Println("Linea 1 SNC - p")
		fmt.Println("Linea 2 SNC - p")

		for turn != 1 {
			// espera
		}

		// Código crítico (SC)
		fmt.Println("Linea 1 SC - p")
		fmt.Println("Linea 2 SC - p")

		turn = 2
	}
}

func q() {
	for {
		// Código SNC
		fmt.Println("Linea 1 SNC - q")
		fmt.Println("Linea 2 SNC - q")
		
		for turn != 2 {
			// espera
		}

		// Código SC
		fmt.Println("Linea 1 SC - q")
		fmt.Println("Linea 2 SC - q")

		turn = 1
	}
}

func main() {
	go p()
	go q()
	time.Sleep(100 * time.Millisecond)
}