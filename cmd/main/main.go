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

	deck.Know = English
	deck.Learning = Russian

	for i := 1; i <= 10; i++ {
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
