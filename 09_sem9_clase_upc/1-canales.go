package main

import "fmt"

func main(){
	// Define el canal sincrono
	mensaje := make(chan string)

	// proceso que emite dato y envía al canal
	go func() {
		mensaje <- "Un nuevo mensaje"
	}()

	msg := <- mensaje	// recibiendo dato del canal

	fmt.Println(msg)
}