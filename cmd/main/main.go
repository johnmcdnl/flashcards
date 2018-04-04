package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	. "github.com/johnmcdnl/flashcards"
)

func main() {

	var deck = NewDeck("deck.db")
	// var deck = NewDeckWithSize("deck.db", 100, 125, false)

	deck.Know = English
	deck.Learning = Russian

	for i := 1; i <= 20; i++ {
		deck.Next()
		deck.Current.PrintQuestion(deck)
		deck.Current.AttemptAnswer(deck.Know, deck.Learning, "")
		deck.SaveState()
		fmt.Println()
		fmt.Println()
	}

	toJSON(deck)

}

func toJSON(i interface{}) {
	j, _ := json.Marshal(i)
	// fmt.Println(string(j))
	ioutil.WriteFile("data.json", j, os.ModePerm)
}
