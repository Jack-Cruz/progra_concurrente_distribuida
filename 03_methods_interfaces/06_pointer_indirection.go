package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// Methods with a pointer receiver taker either a value or a pointer as a receiver
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Functions with a pointer argument must take a pointer
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f;
	v.Y = v.Y * f; 
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)	// -> (&v).Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}