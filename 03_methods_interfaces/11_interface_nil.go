package main

import "fmt"

type Inter interface {
	M()
}

// Text
type Text struct {
	S string
}

func (text *Text) M() {
	if text == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(text.S)
}

// Extra methods
func describe(inter Inter) {
	fmt.Printf("(%v, %T)\n", inter, inter)
}

func main() {
	var elem Inter

	var text *Text
	elem = text
	describe(elem)
	elem.M()

	elem = &Text{"hello"}
	describe(elem)
	elem.M()
}
