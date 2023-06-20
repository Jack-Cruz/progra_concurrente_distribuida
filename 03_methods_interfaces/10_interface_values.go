package main

import (
	"fmt"
	"math"
)
// Interface inter
type Inter interface {
	M()
}

// Text
type Text struct {
	S string
}

func (text *Text) M() {
	fmt.Println(text.S)
}

// Decimal
type Decimal float64

func (decimal Decimal) M() {
	fmt.Println(decimal)
}

// Extra methods
func describe(inter Inter) {
	fmt.Printf("(%v, %T)\n", inter, inter)
}

func main() {
	var elem Inter

	elem = &Text{"Hello"}
	describe(elem)
	elem.M()

	elem = Decimal(math.Pi)
	describe(elem)
	elem.M()

}