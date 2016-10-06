package main

import (
	"fmt"
	// "io"
	"os"
	// "math"
	"errors"
	"github.com/fatih/color"
	"strconv"
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

func (b *TicTacGoBoard) PlaceOnBoard(position int, character string) (err error) {
	if position > (len(b) * len(b[0])) {
		return errors.New("Error placing " + character + " in board: position is too high!")
	}
	if position <= 0 {
		return errors.New("Error placing " + character + " in board: position is too low!")
	}
	currentPosition := 1
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[0]); j++ {
			if position == currentPosition {
				if b[i][j] != "" {
					return errors.New("Error placing " + character + " in board: something already exists there!")
				}
				b[i][j] = character
				currentPosition += 1
				return nil
			} else {
				currentPosition += 1
			}
		}
	}
	return nil
}


// not sure I fully understand this bit.
// first, we create a TicTacGoBoard literal
// then, we return the 
func NewTicTacGoBoard() *TicTacGoBoard {
	return &TicTacGoBoard{{"", "", ""}, {"", "", ""}, {"", "", ""}}
}

func GameLoop() {
	key := TicTacGoBoard{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	gameBoard := NewTicTacGoBoard() // a pointer to a TicTacGoBoard
	input := os.Stdin               //os.Stdin is a pointer to a os.File
	inputBuffer := make([]byte, 100)
	opponent := NewOpponent()
	for {
		fmt.Println("key:")
		fmt.Println(key)
		fmt.Println("board:")
		fmt.Println(gameBoard)
		fmt.Println("Which space would you like to write to?")
		fmt.Printf("Input number here: ")
		n, _ := input.Read(inputBuffer)
		inputString := strings.TrimSpace(string(inputBuffer[:n]))
		i, err := strconv.Atoi(inputString)
		if err != nil {
			fmt.Println("that's not a real number!")
			fmt.Println(err)
		} else {
			fmt.Println(i)
			err := gameBoard.PlaceOnBoard(i, "X")
			if err != nil {
				c := color.New(color.FgRed).Add(color.Underline)
				c.Println(err)
			}
			fmt.Println("opponent playing...")
			opponent.PlaceOnBoard(gameBoard)
		}
	}

}

func main() {

	fmt.Printf("Welcome to Tic Tac Go!\n")
	GameLoop()
	// clear string
	fmt.Printf("Thanks for playing!\n")
}
