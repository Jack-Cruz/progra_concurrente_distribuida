package main

import (
	"fmt"
	"math/rand"
)

const (
	MAX = 9	// nro de coincidencias
	K = 4	// nro de digitos a imprimir antes de un salto de linea
)

func send(inC chan rune) {
	for {
		inC <- rune(rand.Intn(26)+65)	// 65: Valor que representa la letra A (ASCII)
	}
}

func compress(inC, pipe chan rune) {
	n := 0
	previous := <- inC

	for {
		c := <- inC // Segunda lectura
		if c == previous && n < MAX-1 {
			n++
		} else {
			if n > 0 {
				pipe <- rune(n + 49)	// 49: numero 0 (ASCII)
				n = 0
			}
			pipe <- previous
			previous = c
		}
	}
}

func output(pipe, outC chan rune) {
	m := 0
	for {
		// c := <- pipe
		// outC <- c
		outC <- <- pipe

		m++

		// Condición para el salto de linea
		if m >= K {
			outC <- '\n'
			m = 0
		}
	}
}

func main() {
	// creación de los canales a usar
	inC := make(chan rune)
	pipe := make(chan rune)
	outC := make(chan rune)

	// procesos concurrentes
	go send(inC)
	go compress(inC, pipe)
	go output(pipe, outC)

	// proceso que recibe del output
	for {
		fmt.Printf("%c", <-outC)
	}
}