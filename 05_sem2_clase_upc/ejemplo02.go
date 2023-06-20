package main

import (
	"fmt"
	"time"
)

var n int = 0

func p() {
	k1 := 1
	n = k1
}

func q() {
	k2 := 2
	n = k2
}

func main() {
	go p()
	go q()

	time.Sleep(100 * time.Millisecond)
	// Tmprimir el valor final de n
	fmt.Println("El valor final de n es ", n)
}