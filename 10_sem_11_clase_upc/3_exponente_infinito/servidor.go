package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func handle(conn net.Conn){
	defer conn.Close()
	
	//////////////////////////////////////////////////////
	// Recuperar los datos que son enviados por el cliente
	read := bufio.NewReader(conn)

	msg, _ := read.ReadString('\n')
	msg = strings.TrimSpace(msg)	// Quita los saltos de linea
	num, _ := strconv.Atoi(msg)		// Comvierte a entero

	fmt.Println("Llegó el numero:", num)
	
	// Retorno al cliente
	fmt.Fprintln(conn, num*num)

}

func main() {
	// Rol del servidor
	// Escuchar
	ln, error := net.Listen("tcp", "localhost:8000") 
	if error != nil {
		log.Println("Error al escuchar. Detalle ", error.Error())
		os.Exit(1)
	}

	defer ln.Close()

	for {
		// Aceptar la comunicación del cliente
		conn, err := ln.Accept()
		
		if err != nil {
			log.Println("Fallo en la aceptación de la conexión. Detalle ", err.Error())
			// Manejo adicional
		}
		
		go handle(conn)
	}
}