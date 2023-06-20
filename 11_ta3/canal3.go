package main

import "fmt"

// Enviar en un solo flujo de varios nodos con bucle

func enviar(ch chan int, n int){
	ch <- n
}

func procesar(left, right chan int){
	n := <- left
	right <- n
}

func recibir(ch chan int){
	n := <- ch
	fmt.Println(n)
}

func main(){
	col := 7
	ch := make([]chan int, col)
	for i := range ch {
		ch[i] = make(chan int)
	}
	go enviar(ch[0], 3)
	for i := 0; i < col-1; i++ {
		go procesar(ch[i], ch[i+1])
	}
	recibir(ch[6])

	fmt.Println("Hello world")
}