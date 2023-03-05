package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	board[0][0] = "1"
	board[0][1] = "2"
	board[0][2] = "3"
	board[1][0] = "4"
	board[1][1] = "5"
	board[1][2] = "6"
	board[2][0] = "7"
	board[2][1] = "8"
	board[2][2] = "9"
	fmt.Println("--------------")
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
