package main

import "fmt"

type Inter interface {
	M()
}

type Text struct {
	S string
}

// This method defines 'Text' implement 'Inter'
// we don't need to explicitly declare it does so.
func (text Text) M() {
	fmt.Println(text.S)
}

func main() {
	var elem Inter = Text{"hello"}
	elem.M()
}
