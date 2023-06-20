package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main(){
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line numer.
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	// A slice of names
	names := []string{"Gladys", "Samantha", "Darrin", "Jack"}


	// Request a greeting message.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned map of
	// messages to the console
	fmt.Println(messages)
}

// Go code for fun!
package main

import "fmt"

const {
	// Create a huge number by shifting a 1 bit left 100 places
	// In order words. the binary number that is 1 followed by 100 zeroes.
	Big = i << 100
	Small = Big >> 99
}

func needInt(x int) int { return x*10 + 1}
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println("Big  : ", Big)
	fmt.Println("Small: ", )
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
