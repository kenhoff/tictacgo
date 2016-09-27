package main

import (
	"fmt"
	// "io"
	"os"
	// "math"
	"strings"
)

type TicTacGoBoard [3][3]string

func (b TicTacGoBoard) String() string {
	rows := make([]string, 0)
	for i := 0; i < len(b); i++ {
		row := make([]string, 0)
		for j := 0; j < len(b[i]); j++ {
			var char string
			if b[i][j] == "" {
				char = " "
			} else {
				char = b[i][j]
			}
			row = append(row, char)
		}
		rows = append(rows, strings.Join(row, " | ")) // strings.Join only takes a slice, not an array.
	}
	boardString := strings.Join(rows, "\n- + - + -\n")
	return fmt.Sprint(boardString)
}

func (b *TicTacGoBoard) placeOnBoard(position int, character string) {
	currentPosition := 1
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[0]); j++ {
			fmt.Printf("looking at position %v, %v: %v\n", i, j, b[i][j])
			if position == currentPosition {
				b[i][j] = character
				currentPosition += 1
				fmt.Printf("updated: %v\n", b[i][j])
			} else {
				currentPosition += 1
			}
		}
	}
}

func NewTicTacGoBoard() *TicTacGoBoard {
	return &TicTacGoBoard{{"", "", ""}, {"", "", ""}, {"", "", ""}}
}

func main() {
	key := TicTacGoBoard{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	gameBoard := NewTicTacGoBoard()
	input := os.Stdin //os.Stdin is a pointer to a os.File
	inputBuffer := make([]byte, 100)

	fmt.Printf("Welcome to Tic Tac Go!\n")

	fmt.Println("key:")
	fmt.Println(key)
	fmt.Println("board:")
	fmt.Println(gameBoard)
	fmt.Printf("Which space would you like to write to?\nInput number here: ")

	n, _ := input.Read(inputBuffer)
	fmt.Printf("%v bytes read: %v\n", n, string(inputBuffer))
	fmt.Println(inputBuffer)
	gameBoard.placeOnBoard(5, "X")
	fmt.Println(gameBoard)

	// clear string
	fmt.Printf("Thanks for playing!\n")
}
