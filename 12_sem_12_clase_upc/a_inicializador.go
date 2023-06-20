package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var addressRemoto string

func main() {
	// Lectura de consola del host de origen
	brIn := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese el puerto del host remoto: ")
	puertoRemoto, _ := brIn.ReadString('\n')
	puertoRemoto = strings.TrimSpace(puertoRemoto)
	addressRemoto = fmt.Sprintf("localhost:%s", puertoRemoto)

	enviar(6)
	enviar(3)
	enviar(1)
	enviar(5)
}

func enviar(num int){
	conn, _ := net.Dial("tcp", addressRemoto)
	defer conn.Close()

	fmt.Fprintf(conn, "%d\n", num)
}
