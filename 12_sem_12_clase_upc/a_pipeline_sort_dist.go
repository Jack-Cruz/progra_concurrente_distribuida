package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"os"
)

const (
	N = 4
	PORT1 = "8001"	// Origen
	PORT2 = "8002"	// Destino
)

var min int
var cont int
var chCont chan int
var numero int
var addressLocal string
var addressRemoto string

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

	// Lectura del nro de mensajes a recibir
	brIn = bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese el numero de mensajes a recibir: ")
	numStr, _ := brIn.ReadString('\n')
	numStr = strings.TrimSpace(numStr)
	numero, _ = strconv.Atoi(numStr)

	chCont = make(chan int, 1)
	chCont <- 0

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
	dato, _ := br.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(dato))

	// Lógica del ordenamiento
	cont = <- chCont
	if cont == 0 {
		min = num
	} else {
		if num < min {
			enviar(min)
			min = num
		} else {
			enviar(num)
		}
	}
	cont++
	if cont == numero {
		fmt.Printf("Num: %d\n", min)
	}
	chCont <- cont
}

func enviar(num int) {
	conn, _ := net.Dial("tcp", addressRemoto)
	defer conn.Close()

	fmt.Fprintf(conn, "%d\n", num)
}