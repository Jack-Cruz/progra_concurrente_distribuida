package main

import "fmt"

func pi(n int) float64 {
    N := float64(n)

    nth := 4
    res := make(chan float64)

    for i := 0; i < nth; i++ {

        go func(pid int) {
            sum := 0.
            for i := pid; i < n; i += nth {
                x := (float64(i) + .5) / N
                sum += 4. / (1. + x*x)
            }
            res<- sum
        }(i)
    }
    sum := 0.
    for i := 0; i < nth; i++ {
        sum += <-res
    }

    return sum / N
}

func main() {
    fmt.Println(pi(1000000))
}

