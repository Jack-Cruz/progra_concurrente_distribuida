package main

import "fmt"

type Inter interface {
	M()
}

// Extra methods
func describe(inter Inter) {
	fmt.Printf("(%v, %T)\n", inter, inter)
}

func main() {
	var elem Inter
	describe(elem)
	elem.M()
}