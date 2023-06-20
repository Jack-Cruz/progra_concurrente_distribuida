package main

import "fmt"

var end chan bool

// Enviar y recibir problema de condición de carrera

func enviar(ch chan int, n int){
	ch <- n
}

func procesar(left, right chan int){
	n := <- left
	<- right
	// siempre gana el izquierdo
	right <- n
}

func recibir(ch chan int){
	n := <- ch
	fmt.Println(n)
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
	go enviar(ch[0], 1)
	go enviar(ch[1], 2)
	go procesar(ch[0], ch[1])
	go recibir(ch[1])
	
	// Esperar a finalizar
	<- end
	fmt.Println("Hello world")
}