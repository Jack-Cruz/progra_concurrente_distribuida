package main

import (
	"fmt"
	"time"
)

var C [10]int = [10]int{39, 35, 41, 35, 90, 64, 1, 64, 5, 3}
var D [10]int

func p(i int){
	var myNumber, count int
	
	// p1
	myNumber = C[i]
	
	// p2
	count = 0
	for _, v := range(C) {
		if v < myNumber {
			count += 1
		}
	}

	// p3
	for {
		if D[count+1-1] == myNumber {
			count += 1;
		} else {
			D[count+1-1] = myNumber
			break
		}
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go p(i)
	}
	time.Sleep(3 * time.Second)

	for _, v := range D {
		fmt.Println(v)
	}
}