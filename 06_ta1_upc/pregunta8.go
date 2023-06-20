package main

import (
	"fmt"
	"time"
)

var n int = 1

func p() {
	for ; n < 1; {
		n = n + 1
		fmt.Println("n:", n)
	}
	fmt.Println("Finished process p")
}

func q() {
	for ; n >= 0; {
		n = n - 1
		fmt.Println("n:", n)
	}
	fmt.Println("Finished process q")
}

func main() {
	go p()
	go q()
	time.Sleep(3 * time.Second)
	fmt.Println("Final n:", n)
}