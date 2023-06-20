package main

import "fmt"

// Extra methods
func describe(inter interface{}) {
	fmt.Printf("(%v, %T)\n", inter, inter)
}

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}