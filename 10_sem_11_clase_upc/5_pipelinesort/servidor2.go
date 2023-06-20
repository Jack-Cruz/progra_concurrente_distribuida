package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

const (
	N = 3
	PORT1 = "8001"	// Origen
	PORT2 = "8002"	// Destino
)

var min int
var cont int
var chCont chan int
// mismo canal, minimo, contador

func main() {
	chCont = make(chan int, 1)
	chCont <- 0

	// Habilitar el modo escucha (servidor)
	ln, _ := net.Listen("tcp", "localhost:" + PORT1)
	defer ln.Close()

	// Manejo de concurrencia para aceptar conexión de clientes
	for {
		conn, _ := ln.Accept()
		go manejador(conn)
	}
}

func manejador(conn net.Conn) {
	defer conn.Close()
	// variables locales
	conread := bufio.NewReader(conn)
	dato, _ := conread.ReadString('\n')
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
	if cont == N {
		fmt.Printf("Num: %d\n", min)
		cont = 0
	}
	chCont <- cont
}

func enviar(num int) {
	conn, _ := net.Dial("tcp", "localhost:"+PORT2)
	defer conn.Close()

	fmt.Fprintf(conn, "%d\n", num)
}