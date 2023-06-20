package main

import (
	"fmt"
	"time"
)


func pagador(ch chan int){
	x := 0
	for {
		fmt.Println("Listo para enviar")
		ch <- x	
		fmt.Println("Aun no sumo")
		x += 1		
	}
}

func cobrador(ch chan int){
	var y int
	for y = range ch{
		fmt.Println("Consumiendo ", y)
	}
}

func main() {
	ch := make(chan int)
	go pagador(ch)
	time.Sleep(5*time.Second)
	cobrador(ch)
}