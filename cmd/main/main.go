package main

import (
	"fmt"

	"github.com/johnmcdnl/flashcards"
)

func main() {

	var deck = flashcards.NewStandardDeck()
	// var deck = flashcards.NewDeckWithSize(flashcards.DeckName, 100, 125, false)

	deck.Know = flashcards.English
	deck.Learning = flashcards.Russian

	for i := 1; i <= 3; i++ {
		deck.Next()
		fmt.Println(deck.Current.PrintQuestion(deck))
		deck.Current.AttemptAnswer(deck.Know, deck.Learning, "")
		deck.SaveState()
	}
}
