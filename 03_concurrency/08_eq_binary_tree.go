package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch := make(chan int)
	var lista1 []int
	var lista2 []int

	go Walk(t1, ch)
	for i := 0; i < 10; i++ {
		lista1 = append(lista1, <-ch)
	}
	go Walk(t2, ch)
	for i := 0; i < 10; i++ {
		lista2 = append(lista2, <-ch)
	}

	for i := 0; i < 10; i++ {
		if lista1[i] != lista2[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(2), tree.New(1)))
}