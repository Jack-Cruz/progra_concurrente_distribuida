package main

import "fmt"

func fibonacci(ch, ch_quit chan int) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <- ch_quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	ch := make(chan int)
	ch_quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		ch_quit <- 0
	}()

	fibonacci(ch, ch_quit)
}