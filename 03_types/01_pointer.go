package main

import "fmt"

func main() {
	i, j := 2, 2701

	p := &i			// point to i
	fmt.Println(*p)	// read it trough the pointer
	*p = 21			// set i trough the pointer
	fmt.Println(*p)	// see the new value of i

	p = &j			// point to j
	*p = *p / 37	// divide j through the pointer
	fmt.Println(j)	// see the new value of j

}