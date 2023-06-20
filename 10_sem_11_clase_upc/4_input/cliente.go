package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	con, _ := net.Dial("tcp", "localhost:8000")
	defer con.Close()

	consoleIn := bufio.NewReader(os.Stdin)
	conIn := bufio.NewReader(con)

	for {
		fmt.Print("Ingrese un mensaje: ")
		msg, _ := consoleIn.ReadString('\n')

		fmt.Fprint(con, msg)
		res, _ := conIn.ReadString('\n')
		fmt.Printf("Cliente: %s", res)
	}
}