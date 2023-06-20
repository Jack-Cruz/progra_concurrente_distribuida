package main

import "fmt"

// Enviar en un solo flujo

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
	go procesar(ch[0], ch[1])
	recibir(ch[1])

	fmt.Println("Hello world")
}