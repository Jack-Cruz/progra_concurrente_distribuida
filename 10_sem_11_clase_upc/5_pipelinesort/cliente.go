package main

import (
	"fmt"
	"net"
)

func main() {
	enviar(5)
	enviar(6)
	enviar(7)
	enviar(5)
}

func enviar(num int) {
	conn, _ := net.Dial("tcp", "localhost:8000")
	defer conn.Close()

	fmt.Fprintf(conn, "%d\n", num)
}
