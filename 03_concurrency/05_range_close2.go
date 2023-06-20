package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main(){
	c := make(chan int, 4)
	go fibonacci(cap(c), c)

	i, ok := <- c
	fmt.Println(i, ok)
	i, ok = <- c
	fmt.Println(i, ok)	
	i, ok = <- c
	fmt.Println(i, ok)	
	i, ok = <- c
	fmt.Println(i, ok)	
	i, ok = <- c
	fmt.Println(i, ok)
}