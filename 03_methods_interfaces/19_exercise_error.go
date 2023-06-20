package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		err := ErrNegativeSqrt(x)
		return x, err
	}
	
	z := x/2
	for i := 1; i <= 10; i++ {
		// Newton method
		diff := (z*z - x) / (2*z)
		z -= diff
		
		if diff < 0.00001 {
			return z, nil;
		}
	}
	
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
