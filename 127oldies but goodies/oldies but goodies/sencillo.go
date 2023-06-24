package main

import "fmt"

func proceso(ch chan int) {
    for i := 0; i < 2; i++ {
        ch<- i + 1
    }
    close(ch)
}

func main() {
    ch := make(chan int)

    go proceso(ch)
    go proceso(ch)
    go proceso(ch)
    go proceso(ch)
    go proceso(ch)

    for val := range ch {
        fmt.Println(val)
    }
}
