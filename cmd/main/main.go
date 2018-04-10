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

	deck.With(
		f.NewPhrase(
			f.NewTranslation(f.En, "goodbye")).With(
			f.NewTranslation(f.Ru, "POKA").With(f.NewPhonetic(f.En, "poka"))).With(
			f.NewTranslation(f.Fr, "au revouir")),
	)

	logrus.Info(deck)
}
