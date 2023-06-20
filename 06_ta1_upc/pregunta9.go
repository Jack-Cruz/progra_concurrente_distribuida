package main

import (
	"fmt"
	"time"
)

var n int = 1
var flag bool = false

func p() {
	for ; flag == false; {
		fmt.Println("Process p, n:", n, "flag:", flag)
		n = 1 - n
	}
}

func q() {
	for ; flag == false; {
		fmt.Println("Process q, n:", n, "flag:", flag)
		if n == 0 {
			flag = true
		}
	}
}

func main() {
	go p()
	go q()
	time.Sleep(3 * time.Second)
	fmt.Println("n   :", n)
	fmt.Println("flag:", flag)
}