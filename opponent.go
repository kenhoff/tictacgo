package main

import (
	"fmt"
	"math/rand"
)

type Opponent struct {
	difficulty string
	token      string
}

func NewOpponent(opponentToken string) *Opponent {
	return &Opponent{"braindead", opponentToken}
}

func (o Opponent) String() string {
	return fmt.Sprintf("This is an opponent with difficulty \"%v\" and token \"%v\".", o.difficulty, o.token)
}

func (o Opponent) PlaceOnBoard(b *TicTacGoBoard) (err error) {
	// first, get available slots from board
	availablePlaces := b.GetOpenPlaces()
	// then, place at that available slot
	if len(availablePlaces) == 0 {
		return nil
	}
	err = b.PlaceOnBoard(availablePlaces[rand.Intn(len(availablePlaces))], o.token)
	if err != nil {
		return err
	} else {
		return nil
	}

}
