package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

var chCont chan int
var addressLocal string
var addressRemoto string


// {"mensaje": 4}

func main() {
	// Lectura de consola del host de origen
	brIn := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese el puerto del host local: ")
	puertoLocal, _ := brIn.ReadString('\n')
	puertoLocal = strings.TrimSpace(puertoLocal)
	addressLocal = fmt.Sprintf("localhost:%s", puertoLocal)

	// Lectura por consola del host de destino
	brIn = bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese el puerto del host remoto: ")
	puertoRemoto, _ := brIn.ReadString('\n')
	puertoRemoto = strings.TrimSpace(puertoRemoto)
	addressRemoto = fmt.Sprintf("localhost:%s", puertoRemoto)

	// Habilitar el modo escucha (servidor) nodo local
	ln, _ := net.Listen("tcp", addressLocal)
	defer ln.Close()
	// Manejo de concurrencia para aceptar conexión de clientes
	for {
		conn, _ := ln.Accept()
		go manejador(conn)
	}
}

func manejador(conn net.Conn) {
	defer conn.Close()
	br := bufio.NewReader(conn)
	str, _ := br.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(str))

	fmt.Printf("Nos ha llegado el %d\n", num)
	if num == 0 {
		fmt.Println("Perdimos! :(")
	} else {
		enviar(num - 1)
	}
}

func enviar(num int) {
	conn, _ := net.Dial("tcp", addressRemoto)
	defer conn.Close()
	fmt.Fprintf(conn, "%d\n", num) 
}