package main

import (
	"bufio"
	"fmt"
	"net"
)

func handle(con net.Conn) {
	defer con.Close()
	
	conIn := bufio.NewReader(con)
	for {
		msg, _ := conIn.ReadString('\n')
		fmt.Printf("Servidor: %s", msg)
		fmt.Fprintln(con, "msg")
	}
	
}

func main() {
	ln, _ := net.Listen("tcp", "localhost:8000")
	defer ln.Close()
	for {
		con, _ := ln.Accept()
		go handle(con)
	}
}
