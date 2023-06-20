package main

import (
	"fmt"
	"math"
)

// MyFloat

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f > 0 {
		return float64(-f)
	}
	return float64(f)
}

// Vertex
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// Interface Abser
type Abser interface {
	Abs() float64
}

func main(){
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f 	// a MyFloat implements Abser
	a = &v	// a *Vertex implements Abser

	// Here, v is a Vetex (not *Vertex) so it not implement Abser
	a = &v
	fmt.Println(a.Abs())

}