package main

import "fmt"

// Enviar a la derecha en un mismo nodo, problema de condicion de carrera

var end chan bool

func enviar(ch chan int, n int){
	ch <- n
}

func procesar(left, right chan int){
	n := <- left
	<- right
	// siempre gana el izquierdo
	right <- n
}

func derecha(left, right chan int){
	n := <- left
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
	go enviar(ch[1], 1)
	go enviar(ch[2], 2)
	go procesar(ch[1], ch[2])
	go derecha(ch[2], ch[3])
	go recibir(ch[3])
	
	// Esperar a finalizar
	<- end
	fmt.Println("Hello world")
}