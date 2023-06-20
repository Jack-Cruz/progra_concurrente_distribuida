package main

import (
	"fmt"
	"time"
)

// 4 Canales en un nodo a procesar
// 2 entrada
// 2 salida
// envio cualquiera de los dos

func initial(ch chan int, n int){
	ch <- n
}

func enviar(grupo string, ch chan int, n int){
	ch <- n
	fmt.Println("Grupo ", grupo, "envio: ", n)
	for {
		ch <- 0
		fmt.Println("Grupo ", grupo, "envio: ", 0)
	}
}

func procesar(pos1, pos2 string, derecha1, derecha2, izquierda1, izquierda2 chan int){
	var val_pos1, val_pos2 int
	for {
		select {
		case val_pos1 = <- derecha1:
			val_pos2 = <- izquierda2

		case val_pos2 = <- izquierda2:
			val_pos1 = <- derecha1
		}
		fmt.Println("canal1", pos1, "canal2: ", pos2)
		fmt.Println("valor1", val_pos1, "valor2: ", val_pos2)

		if val_pos1 > val_pos2{
			// val_pos1 continua el mismo
			val_pos2 = 0
		}
		if val_pos2 > val_pos1{
			// val_pos2 continua el mismo
			val_pos1 = 0
		}
		select {
		case izquierda1 <- val_pos2:
			derecha2 <- val_pos1
		case derecha2 <- val_pos1:
			izquierda1 <- val_pos2
		}
	}
}

func recibir(grupo string, ch chan int){
	for {
		n := <- ch
		fmt.Println("Grupo ", grupo, "recibi: ", n)
	}
}

func main(){
	col := 7
	ch_right := make([]chan int, col)
	ch_left := make([]chan int, col)
	
	// create channels
	for i := range ch_right {
		ch_right[i] = make(chan int)
	}
	for i := range ch_left {
		ch_left[i] = make(chan int)
	}

	// goroutines
	go enviar("1", ch_right[0], 1)	// Grupo, canal, valor
	go enviar("2", ch_left[2], 2) // Grupo, canal, valor
	//ch_right[1] <- 0
	go initial(ch_left[1], 0)
	go procesar("0", "1", ch_right[0], ch_right[1], ch_left[0], ch_left[1])	// pos1, pos2, derecha1, derecha2, izquierda1, izquierda2 
	go procesar("1", "2", ch_right[1], ch_right[2], ch_left[1], ch_left[2])	// pos1, pos2, derecha1, derecha2, izquierda1, izquierda2 
	go recibir("1", ch_left[0])
	go recibir("2", ch_right[2])
	
	time.Sleep(10 * time.Second)
	fmt.Println("Hello world")
}