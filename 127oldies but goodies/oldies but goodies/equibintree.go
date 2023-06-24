package main

import "fmt"

type Node struct {
	Left   *Node
	Value	int
	Right  *Node
}
func add(node *Node, val int) *Node {
	if node == nil {
		return &Node{Value: val}
	} else if val > node.Value {
		node.Right = add(node.Right, val)
	} else {
		node.Left = add(node.Left, val)
	}
	return node
}
func inOrder(node *Node, f func(int)) {
	if node != nil {
		inOrder(node.Left, f)
		f(node.Value)
		inOrder(node.Right, f)
	}
}
type Tree struct {
	root   *Node
	length	int
}
func (t *Tree)Add(val int) {
	t.root = add(t.root, val)
	t.length++
}
func (t *Tree)InOrder(f func(int)) {
	inOrder(t.root, f)
}
func check(t1, t2 *Tree) bool {
	if t1.length != t2.length {
		return false
	}
	ch1 := make(chan int)
	ch2 := make(chan int)
	go t1.InOrder(func(a int) { ch1<- a })
	go t2.InOrder(func(a int) { ch2<- a })
	var a, b int
	for i := 0; i < t1.length; i++ {
		select {
		case a = <-ch1:
			b = <-ch2
		case b = <-ch2:
			a = <-ch1 
		}
		if a != b {
			return false
		}
	}
	return true
}
func main() {
	n := 10
	t1 := new(Tree)
	t2 := new(Tree)
	for i := 0; i < n; i++ {
		t1.Add(i + 2)
		t2.Add(10 - i)
	}
	print := func(a int) {
		fmt.Print(a, " ")
	}
	t1.InOrder(print)
	fmt.Println()
	t2.InOrder(print)
	fmt.Println()
	if check(t1, t2) {
		fmt.Println("Arboles son equivalentes :)")
	} else {
		fmt.Println("Arboles NO son equivalentes :(")
	}
}