package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello, Reader!")
	result := make([]byte, 8)

	for {
		n, err := reader.Read(result)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, result)
		fmt.Printf("b[:n] = %q", result[:n])
		if err == io.EOF {
			break
		}
	}
}