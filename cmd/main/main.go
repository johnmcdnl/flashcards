package main

import "github.com/johnmcdnl/flashcards"

func main() {

	// var deck = NewDeck("deck.db")
	var deck = flashcards.NewDeckWithSize("deck.db", 100, 125, false)

	deck.Know = flashcards.English
	deck.Learning = flashcards.Russian

	for i := 1; i <= 3; i++ {
		deck.Next()
		deck.Current.PrintQuestion(deck)
		deck.Current.AttemptAnswer(deck.Know, deck.Learning, "")
		deck.SaveState()
	}
}
