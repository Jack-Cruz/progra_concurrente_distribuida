package main

import "fmt"

func main() {
	pow := make([]int, 10)

	for i := range pow {		// Omit the values
		pow[i] = 1 << uint(i)	// == 2**i
	}
	for _, value := range pow {	// Omit the indices
		fmt.Printf("%d\n", value)
	}
}

