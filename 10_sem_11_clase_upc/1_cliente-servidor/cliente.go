package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Rol del cliente
	// Envia datos
	conn, _ := net.Dial("tcp", "localhost:8000")
	defer conn.Close()

	// Enviar datos al servidor
	fmt.Fprintln(conn, "enviado datos del cliente Jack!!")

	// Recibe lo que le llega del servidor
	read := bufio.NewReader(conn)

	msg, _ := read.ReadString('\n')
	fmt.Println(msg)
	
	
}