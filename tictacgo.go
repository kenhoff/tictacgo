package main

import (
	"fmt"
	"io"
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

func NewTicTacGoBoard() *TicTacGoBoard {
	return &TicTacGoBoard{{"", "", ""}, {"", "", ""}, {"", "", ""}}
}

func main() {
	key := TicTacGoBoard{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	gameBoard := NewTicTacGoBoard()
	input := os.Stdin //os.Stdin is a pointer to a os.File
	inputBuffer := make([]byte, 100)

	fmt.Printf("Welcome to Tic Tac Go!\n")

	for { // game loop
		fmt.Println("key:")
		fmt.Println(key)
		fmt.Println("board:")
		fmt.Println(gameBoard)
		fmt.Printf("Which space would you like to write to?\nInput number here: ")

		n, err := input.Read(inputBuffer)
		if err == io.EOF {
			break
			fmt.Println("Done reading input")
		} else if err != nil {
			fmt.Errorf("Error! ", err)
		}
		fmt.Printf("%v bytes read: %v\n", n, string(inputBuffer))
		// clear string
	}
	fmt.Printf("Thanks for playing!\n")
}
