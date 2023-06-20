package main

import (
	"fmt"
)

func main() {
	// Create a tic-tac-toe board
	board := [][]string {
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players tale turns
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	board[0][1] = "X"
	board[1][1] = "O"

	// One way to display
	// for i := 0; i < len(board); i++ {
	// 	fmt.Printf("%s\n", strings.Join(board[i], " "))
	// }

	// Second way to display
	for i := 0; i < len(board); i++ {
		row := board[i]
		for j := 0; j < len(row); j++ {
			fmt.Printf("%s ", row[j])
		}
		fmt.Println()
	}
}