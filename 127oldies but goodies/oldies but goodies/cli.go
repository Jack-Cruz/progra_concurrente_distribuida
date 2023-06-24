package main

import (
    "net"
    "fmt"
)

func main() {
    con, err := net.Dial("tcp", "localhost:8000")
    if err != nil {
        fmt.Println("No se pudo conectar.")
        return
    }
    defer con.Close()
    fmt.Fprintln(con, "Hola, soy cliente!")
}
