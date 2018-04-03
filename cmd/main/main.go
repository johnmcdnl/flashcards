package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/johnmcdnl/flashcards"
)

func main() {
	deck := flashcards.NewDeck()

	toJson(deck)
}

func toJson(i interface{}) {
	j, _ := json.Marshal(i)
	fmt.Println(string(j))
	ioutil.WriteFile("data.json", j, os.ModePerm)
}
