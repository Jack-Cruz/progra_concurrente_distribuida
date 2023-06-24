package main

import "fmt"

func source(r []int, s chan int) {
    for _, v := range r {
        s<- v
    }
    close(s)
}
func zero(n int, w chan int) {
    for i := 0; i < n; i++ {
        w<- 0
    }
    close(w)
}
func sink(n chan int) {
    for _ = range n {}
}
func result(c [][]int, i int, e chan int) {
    j := 0
    for v := range e {
        c[i][j] = v
        j++
    }
    fin<- true
}
func multiplier(fE int, n, e, s, w chan int) {
    for sE := range n {
        sum := <-e
        sum = sum + fE * sE
        s<- sE
        w<- sum
    }
    close(s)
    close(w)
}
var fin chan bool
func main() {
    a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
    b := [][]int{{1, 0, 2}, {0, 1, 2}, {1, 0, 0}}
    c := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
    nra := len(a)
    nca := len(a[0])
    fin = make(chan bool)
    h := make([][]chan int, nra)
    for i := range h {
        h[i] = make([]chan int, nca+1)
        for j := range h[i] {
            h[i][j] = make(chan int)
        }
    }
    v := make([][]chan int, nra+1)
    for i := range v {
        v[i] = make([]chan int, nca)
        for j := range v[i] {
            v[i][j] = make(chan int)
        }
    }
    for i := range a { // nra iteraciones
        go zero(nca, h[i][nca])
        go result(c, i, h[i][0])
    }
    for i, row := range b { // nca iteraciones
        go source(row, v[0][i])
        go sink(v[nra][i])
    }
    for i, row := range a {
        for j, val := range row {
            go multiplier(val, v[i][j], h[i][j+1],
                               v[i+1][j], h[i][j])
        }
    }
    for _ = range a {
        <-fin
    }
    fmt.Println(c)
}
