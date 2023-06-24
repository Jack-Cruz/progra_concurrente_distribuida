package main

import (
    "fmt"
    "time"
)

func fork(ch chan bool) {
    for {
        ch<- true
        <-ch
    }
}
func philosopher(i int, right, left chan bool) {
    for {
        fmt.Printf("Phil%d thinking\n", i)
        <-right
        <-left
        fmt.Printf("Phil%d eating\n", i)
        right<- true
        left<- true
    }
}

func main() {
    n       := 5
    forks   := make([]chan bool, n)
    for i := range forks {
        forks[i] = make(chan bool, 1)
    }
    for i, f := range forks {
        go fork(f)
        go philosopher(i, forks[i],
                          forks[(i+1)%n])
    }
    time.Sleep(time.Second*10)
}
