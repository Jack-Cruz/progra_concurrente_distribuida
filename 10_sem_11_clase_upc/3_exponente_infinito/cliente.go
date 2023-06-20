package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"time"
)

// Envia datos
func enviar_recibir(num int) {
	// enviar
	conn, _ := net.Dial("tcp", "localhost:8000")
	defer conn.Close()
	fmt.Fprintln(conn, num)

	// recibir
	read := bufio.NewReader(conn)
	msg, _ := read.ReadString('\n')
	fmt.Println(msg)

}
func main() {
	// Enviar datos al servidor
	rand.Seed(time.Now().Unix())

	for i := 0; i < 10; i++ {
		// num := rand.Intn(100)
		enviar_recibir(i)
	}
}