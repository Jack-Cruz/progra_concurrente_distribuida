package main

import (
    "fmt"
    "net"
    "bufio"
)

func handle(con net.Conn) {
    defer con.Close()
    r := bufio.NewReader(con)
    for {
        msg, _ := r.ReadString('\n')
        fmt.Fprint(con, msg)

        fmt.Println("recibido: ", msg)
        if len(msg) == 0 || msg[0] == 'x' {
            break
        }
    }
}

func main() {
    ln, _ := net.Listen("tcp", "10.11.98.229:8000")
    defer ln.Close()
    for {
        con, _ := ln.Accept()
        go handle(con)
    }
}
