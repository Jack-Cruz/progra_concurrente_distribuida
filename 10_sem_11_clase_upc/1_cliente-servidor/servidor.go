package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// Rol del servidor
	// Escuchar
	ln, error := net.Listen("tcp", "localhost:8000") 
	if error != nil {
		log.Println("Error al escuchar. Detalle ", error.Error())
		os.Exit(1)
	}

	defer ln.Accept()

	// Aceptar la comunicación del cliente
	conn, err := ln.Accept()
	
	if err != nil {
		log.Println("Fallo en la aceptación de la conexión. Detalle ", err.Error())
		// Manejo adicional
	}
	defer conn.Close()

	//////////////////////////////////////////////////////
	// Recuperar los datos que son enviados por el cliente
	read := bufio.NewReader(conn)

	msg, _ := read.ReadString('\n')
	
	fmt.Println(msg)
}