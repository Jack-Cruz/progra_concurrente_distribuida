package main

import "fmt"

func main() {
	m := make(map[string]int)

	// Insert or update
	m["Answer"] = 42
	fmt.Println("The value: ", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value: ", m["Answer"])

	// delete
	delete(m, "Answer")
	delete(m, "Woops!")

	// test that key is present
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}