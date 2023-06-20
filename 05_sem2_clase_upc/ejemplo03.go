package main

import (
	"fmt"
	"time"
)

var n  = 0

func proceso(){
	var temp int

	temp = n
	n = temp + 1
	time.Sleep(100 * time.Nanosecond)
}

func main() {
	go proceso();	// proceso 0
	go proceso();	// proceso q

	time.Sleep(100 * time.Millisecond)
	fmt.Println("El valor de n es = ", n)
}