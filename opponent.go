package main

import (
	"fmt"
)

type Opponent struct {
	difficulty string
	token      string
}

func NewOpponent() *Opponent {
	return &Opponent{"braindead", "O"}
}

func (o Opponent) String() string {
	return fmt.Sprintf("This is an opponent with difficulty \"%v\" and token \"%v\".", o.difficulty, o.token)
}

func (o Opponent) PlaceOnBoard(b *TicTacGoBoard) (err error) {
	b.PlaceOnBoard(1, o.token)
	if err != nil {
		return err
	} else {
		return nil
	}
}
