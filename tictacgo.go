package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strconv"
	"strings"
)

func GameLoop(playerToken string, opponentToken string) {
	key := TicTacGoBoard{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	gameBoard := NewTicTacGoBoard() // a pointer to a TicTacGoBoard
	input := os.Stdin               //os.Stdin is a pointer to a os.File
	inputBuffer := make([]byte, 100)
	opponent := NewOpponent(opponentToken)
	fmt.Println("key:")
	fmt.Println(key)
	for {
		fmt.Println("board:")
		fmt.Println(gameBoard)
		fmt.Printf("Available places are: %v\n", gameBoard.GetOpenPlaces())
		fmt.Printf("Player %v, which space would you like to take? ", playerToken)
		n, _ := input.Read(inputBuffer)
		inputString := strings.TrimSpace(string(inputBuffer[:n]))
		i, err := strconv.Atoi(inputString)
		if err != nil {
			fmt.Println("that's not a real number!")
			fmt.Println(err)
		} else {
			fmt.Println(i)
			err := gameBoard.PlaceOnBoard(i, playerToken)
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
	var playerToken, opponentToken string
	input := os.Stdin //os.Stdin is a pointer to a os.File
	inputBuffer := make([]byte, 100)
	fmt.Println("Welcome to Tic Tac Go!")
	for {
		fmt.Printf("Would you like to be X or O? ")
		n, _ := input.Read(inputBuffer)
		inputString := strings.TrimSpace(string(inputBuffer[:n]))
		if (len(inputString) == 1) && ((inputString == "X") || (inputString == "x") || (inputString == "O") || (inputString == "o")) {
			if (inputString == "X") || (inputString == "x") {
				playerToken = "X"
				opponentToken = "O"
			} else {
				playerToken = "O"
				opponentToken = "X"
			}
			fmt.Printf("Okay! Good luck, %v\n", playerToken)
			break
		} else {
			fmt.Println("That's not an X or an O. Try again, numbnuts.")
		}
	}
	GameLoop(playerToken, opponentToken)
	// clear string
	fmt.Printf("Thanks for playing!\n")
}
