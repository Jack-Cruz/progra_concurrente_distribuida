package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x/2
	for i := 1; i <= 10; i++ {
		// Newton method
		diff := (z*z - x) / (2*z)
		z -= diff
		fmt.Println(z)
		
		if diff < 0.00001 {
			return z;
		}
	}
	return z
}

func main() {
	fmt.Println("Func: ", Sqrt(13))	
	fmt.Println("Math: ", math.Sqrt(13))

}