package main

import "fmt"

func main() {
	board := [][]string{
		{"_", "_", "_"}, // fill with ABC
		{"_", "_", "_"}, // fill with def
		{"_", "_", "_"}, // fill with GHI
	}

	board[0][0] = "A"
	board[0][1] = "B"
	board[0][2] = "C"

	board[1][0] = "d"
	board[1][1] = "e"
	board[1][2] = "f"

	board[2][0] = "G"
	board[2][1] = "H"
	board[2][2] = "I"

	for x := 0; x < len(board); x++ {
		fmt.Println(board[x])
	}
}
