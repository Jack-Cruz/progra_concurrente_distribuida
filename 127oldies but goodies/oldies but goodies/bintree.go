package main

import "fmt"

type Node struct {
    Left    *Node
    Value   int
    Right   *Node
}

func inOrder(n *Node, ch chan int) {
    if n != nil {
        inOrder(n.Left, ch)
        ch<- n.Value
        inOrder(n.Right, ch)
    }
}

func add(n *Node, e int) *Node {
    if n == nil {
        return &Node{Value: e}
    } else if e > n.Value {
        n.Right = add(n.Right, e)
    } else {
        n.Left = add(n.Left, e)
    }
    return n
}

func check(r1, r2 *Node) bool {
    ch1 := make(chan int)
    ch2 := make(chan int)
    go inOrder(r1, ch1)
    go inOrder(r2, ch2)
    for i := 0; i < 5; i++ {
        if <-ch1 != <-ch2 {
            return false
        }
    }
    return true
}

func main() {
    var r1 *Node
    r1 = add(r1, 5)
    r1 = add(r1, 3)
    r1 = add(r1, 4)
    r1 = add(r1, 1)
    r1 = add(r1, 2)
    var r2 *Node
    r2 = add(r2, 1)
    r2 = add(r2, 2)
    r2 = add(r2, 3)
    r2 = add(r2, 4)
    r2 = add(r2, 5)
    if check(r1, r2) {
        fmt.Println(":)")
    } else {
        fmt.Println(":(")
    }
}
