package main

import (
    "fmt"
    "net"
    "bufio"
)

func handle(con net.Conn, id int) {
    defer con.Close()
    r := bufio.NewReader(con)
    for {
        msg, _ := r.ReadString('\n')
        fmt.Fprint(con, msg)
        fmt.Printf("Con%d: %s", id, msg)
        if len(msg) == 0 || msg[0] == 'x' {
            break
        }
    }
    fmt.Printf("Con%d cerrada!\n", id)
}

func main() {
    ln, _ := net.Listen("tcp", "10.142.112.54:8000")
    defer ln.Close()
    cont := 0
    for {
        con, _ := ln.Accept()
        go handle(con, cont)
        cont++
    }
}
