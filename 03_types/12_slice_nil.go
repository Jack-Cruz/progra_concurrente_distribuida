package main

import "fmt"

func main() {
	var s []int			// zero value of a slice
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}