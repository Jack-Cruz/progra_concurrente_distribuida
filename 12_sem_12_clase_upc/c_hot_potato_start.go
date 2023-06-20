package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

var addressRemoto string

func main() {
	// Lectura de consola del host de origen
	br := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese el puerto del host local: ")
	puertoRemoto, _ := br.ReadString('\n')
	puertoRemoto = strings.TrimSpace(puertoRemoto)
	addressRemoto = fmt.Sprintf("localhost:%s", puertoRemoto)
	for {
		fmt.Print("Ingrese un numero: ")
		str, _ := br.ReadString('\n')
		num, _ := strconv.Atoi(strings.TrimSpace(str))
		enviar(num)
	}
}

func enviar(num int) {
	conn, _ := net.Dial("tcp", addressRemoto)
	defer conn.Close()
	fmt.Fprintf(conn, "%d\n", num) 
}