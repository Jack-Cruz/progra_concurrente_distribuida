package main

import (
	"fmt"
	"time"
)

var n int = 0

func p(K int) {
	var temp int
	for i := 1; i <= K; i++ {
		temp = n
		n = temp + 1
		fmt.Println("proceso p: ", n)
	}
}

func q(K int) {
	var temp int
	for i := 1; i <= K; i++ {
		temp = n
		n = temp - 1
		fmt.Println("proceso q: ", n)
	}
}

func main() {
	// goroutines (proceso concurrente)
	K := 2
	go p(K)
	go q(K)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Valor final de n:", n)
}