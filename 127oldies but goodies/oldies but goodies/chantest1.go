package main

import (
    "fmt"
)

func main() {
    var ch1 chan int
    ch1 = make(chan int)

    go func() {
        ch1<- 10
    }()

    // NO PODEMOS ENVIAR Y RECIBIR DE UN CANAL
    // SINCRONO EN EL MISMO PROCESO!
    a := <-ch1

    fmt.Printf("Bye %d\n", a)
}
