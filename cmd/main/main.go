package main

import (
	"fmt"

	f "github.com/johnmcdnl/flashcards"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)
	fmt.Println("It runs")
	var deck = f.NewDeck(nil)

	for i := 0; i < 10; i++ {
		phrase := deck.Next()
		fmt.Println(phrase)
		logrus.Info(phrase.Get(f.Ru))
		logrus.Errorln(deck.Answer(phrase, f.En, "hello"))
	}

}
