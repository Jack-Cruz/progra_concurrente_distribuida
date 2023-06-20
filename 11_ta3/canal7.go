package main

import (
	"fmt"
	"time"
)

// 4 Canales en un nodo a procesar
// 2 entrada
// 2 salida

var end chan bool

func enviar(grupo string, ch chan int, n int){
	ch <- n
	fmt.Println("Grupo ", grupo, "envio: ", n)
	for {
		ch <- 0
		fmt.Println("Grupo ", grupo, "envio: ", 0)
	}
}

func procesar(pos1, pos2 string, derecha1, derecha2, izquierda1, izquierda2 chan int){
	val_pos1 := <- derecha1
	val_pos2 := <- izquierda2
	fmt.Println("canal1", pos1, "canal2: ", pos2)

	if val_pos1 > val_pos2{
		// val_pos1 continua el mismo
		val_pos2 = 0
	}
	if val_pos2 > val_pos1{
		// val_pos2 continua el mismo
		val_pos1 = 0
	}

	derecha2 <- val_pos1
	izquierda1 <- val_pos2
}

func recibir(grupo string, ch chan int){
	n := <- ch
	fmt.Println("Grupo ", grupo, "recibi: ", n)
	end <- true
}

func main(){
	col := 7
	ch_right := make([]chan int, col)
	ch_left := make([]chan int, col)
	
	// create channels
	end = make(chan bool)
	for i := range ch_right {
		ch_right[i] = make(chan int)
	}
	for i := range ch_left {
		ch_left[i] = make(chan int)
	}

	// goroutines
	go enviar("1", ch_right[0], 5)	// Grupo, canal, valor
	go enviar("2", ch_left[1], 8) // Grupo, canal, valor
	go procesar("0", "1", ch_right[0], ch_right[1], ch_left[0], ch_left[1])	// pos1, pos2, derecha1, derecha2, izquierda1, izquierda2 
	go recibir("1", ch_left[0])
	go recibir("2", ch_right[1])
	
	time.Sleep(10 * time.Second)
	fmt.Println("Hello world")
}