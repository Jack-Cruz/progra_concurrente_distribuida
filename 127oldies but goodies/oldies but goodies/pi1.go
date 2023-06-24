package main

import "fmt"

func pi(n int) float64 {
    sum := 0.
    N := float64(n)

    for i := 0; i < n; i++ {
        x := (float64(i) + .5) / N
        sum += 4. / (1 + x*x)
    }
    return sum / N
}

func main() {
    fmt.Println(pi(1000000))
}
