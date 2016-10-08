package main

import (
	"errors"
	"fmt"
	"strings"
)

type TicTacGoBoard [3][3]string

// not sure I fully understand this bit.
// first, we create a TicTacGoBoard literal
// then, we return the...?
func NewTicTacGoBoard() *TicTacGoBoard {
	return &TicTacGoBoard{{"", "", ""}, {"", "", ""}, {"", "", ""}}
}

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

func (b *TicTacGoBoard) GetOpenPlaces() []int {
	availablePlaces := []int{}
	currentPosition := 1
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[0]); j++ {
			if b[i][j] == "" {
				availablePlaces = append(availablePlaces, currentPosition)
			}
			currentPosition += 1
		}
	}
	return availablePlaces
}

func (b *TicTacGoBoard) GetBoardWinner() (string, error) {
	for i := 0; i < len(b); i++ { // rows
		rowToken := b[i][0]
		fmt.Println("row token:" + rowToken)
		if rowToken == "" { // if the first row token is blank, then nobody's obviously won that row - move on to the next row
			fmt.Println("breaking")
			continue
		}
		for j := 0; j < len(b[0]); j++ { // else, loop through all the columns
			if b[i][j] != rowToken { // if it's not the same token as the first token in the row, break and move onto the next row
				break
			}
			// if it's the same token, continue on
			if j == (len(b[i]) - 1) { // if it's the end of the row, then the rowToken has won, and return that
				return rowToken, nil
			}
		}
	}
	fmt.Println("returning nil?")
	return "", errors.New("No winner yet!")
}
