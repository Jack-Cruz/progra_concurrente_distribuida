package main

import "fmt"


func pagador(ch chan int){
	for x := 0; x < 5; x++{
		ch <- x			
	}
	close(ch)
}

func cobrador(ch chan int){
	for y := range ch{
		fmt.Println("Consumiendo ", y)
	}
}

func main() {
	ch := make(chan int)
	go pagador(ch)
	cobrador(ch)
}