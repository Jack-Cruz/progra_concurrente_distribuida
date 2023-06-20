package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v *Vertex) Scale1(f float64) {	// Pointer receiver
	v.X = v.X * f;
	v.Y = v.Y * f;
}

func (v Vertex) Scale2(f float64) {	// Value receiver
	v.X = v.X * f;
	v.Y = v.Y * f;
}

func main() {
	v := Vertex{3, 4}
	v.Scale2(10)
	fmt.Println(v.Abs())
}