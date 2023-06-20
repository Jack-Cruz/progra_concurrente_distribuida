package main

import "fmt"

// Enviar en un solo flujo de varios nodos

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
	go procesar(ch[1], ch[2])
	go procesar(ch[2], ch[3])
	go procesar(ch[3], ch[4])
	go procesar(ch[4], ch[5])
	go procesar(ch[5], ch[6])
	recibir(ch[6])

	fmt.Println("Hello world")
}