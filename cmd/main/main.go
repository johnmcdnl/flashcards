package main

import (
	f "github.com/johnmcdnl/flashcards"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	var deck = f.NewDeck(nil)

	deck.With(
		f.NewPhrase(
			f.NewTranslation(f.En, "hello")).With(
			f.NewTranslation(f.Ru, "PRIVET").With(f.NewPhonetic(f.En, "privet"))).With(
			f.NewTranslation(f.Fr, "bonjour")),
	)

	logrus.Info(deck)
}
