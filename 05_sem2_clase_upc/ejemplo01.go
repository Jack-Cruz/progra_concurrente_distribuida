package main

import (
	"fmt"
	"time"
)

func mensaje(i int) {
	fmt.Println("fmt:", i)
}

func main() {
	// golotines (permite lanzar procesos concurrentes)
	for i:=1; i<=6; i++ {
		go mensaje(i)	// oroceso concurrente
	}

	// Instruccion para hacer la pausa
	time.Sleep(100 * time.Millisecond)

}