package main

import ("fmt"
	"time"
	"math/rand"
)

func proc(ch chan int) {
	for {
		dur := time.Duration(rand.Intn(50) + 100)
		time.Sleep(dur)
		ch <- 0
	}
}


func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go proc(ch1)
	go proc(ch2)
	for {
		select {
		case <- ch1:
			fmt.Println("Leido el canal 1 primero")
			<-ch2
		case <- ch2:
			fmt.Println("Leido el canal 2 primero")
			<- ch1
		}
	}
}