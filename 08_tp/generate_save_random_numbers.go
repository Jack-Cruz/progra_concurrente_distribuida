package main

import (
    "fmt"
    "os"
	"math/rand"
)

func main() {
	// List of N numbers
	N := 1000000
	
	// Create a slice of 1,000,000 numbers
    list := make([]int, N)
	
	for i := 0; i < N; i++ {
		list[i] = rand.Intn(N)
	}

    // Open a file for writing
    file, err := os.Create("01_one_million_random_numbers.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    // Write each number to the file on a separate line
    for _, num := range list {
        fmt.Fprintf(file, "%d\n", num)
    }

    fmt.Println("Numbers saved to file.")
}
