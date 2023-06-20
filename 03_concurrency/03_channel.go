package main

import "fmt"

func sum(s []int, c chan int){
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum	// send sum to c
}

func main() {
	s := []int{2, 7, 8, -9, 4, 0}
	
	c := make(chan int)
	
	go sum(s[len(s)/2:], c)
	go sum(s[:len(s)/2], c)

	x, y := <-c, <-c// receive from c (wait current routine until data is received from channel)
	
	fmt.Println(x, y, x+y)
}