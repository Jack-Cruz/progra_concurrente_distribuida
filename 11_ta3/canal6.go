package main

import (
	"fmt"
	"time"
)

// enviar, recibir, procesar en un mismo nodo

var end chan bool

func enviar(grupo string, ch chan int, n int){
	ch <- n
	fmt.Println("Grupo ", grupo, "envio: ", n)
	for {
		ch <- 0
		fmt.Println("Grupo ", grupo, "envio: ", 0)
	}
}

func procesar(idleft, idright string, ch_left, ch_right chan int){
	val_left := <- ch_left
	val_right := <- ch_right
	fmt.Println("Procesar: ")
	fmt.Println("canal ", idleft, "recibi: ", val_left)
	fmt.Println("canal ", idright, "recibi: ", val_right)

	if val_left > val_right{
		// val_left continua el mismo
		val_right = 0
	}
	if val_right > val_left{
		// val_right continua el mismo
		val_left = 0
	}

	ch_right <- val_left
	ch_left <- val_right
}

func recibir(grupo string, ch chan int){
	n := <- ch
	fmt.Println("Grupo ", grupo, "recibi: ", n)
	end <- true
}

func main(){
	col := 7
	ch := make([]chan int, col)
	
	// create channels
	end = make(chan bool)
	for i := range ch {
		ch[i] = make(chan int)
	}

	// goroutines
	
	go enviar("1", ch[1], 1)
	go enviar("2", ch[2], 2)
	go procesar("0", "1", ch[0], ch[1])
	go procesar("1", "2", ch[1], ch[2])
	go procesar("2", "3", ch[2], ch[3])
	go recibir("1", ch[0])
	go recibir("2", ch[3])
	
	time.Sleep(10 * time.Second)
	fmt.Println("Hello world")
}