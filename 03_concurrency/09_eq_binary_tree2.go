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
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)	
	go Walk(t2, ch2)
	
	var g1 []int {}
	var g2 []int {}
	for i := range ch1 {
		append(g1, i)
	}
	for i := range ch2 {
		append(g2, i)
	}
	return g1[len(g1)-1] == g2[len(g2)-1]
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(2)))
}