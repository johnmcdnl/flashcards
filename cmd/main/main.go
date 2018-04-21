package main

import (
	f "github.com/johnmcdnl/flashcards"
	"github.com/sirupsen/logrus"
	"fmt"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)
	fmt.Println("It runs")
	var deck = f.NewDeck(nil)

	//deck.With(
	//	f.NewPhrase(
	//		f.NewTranslation(f.En, "hello")).With(
	//		f.NewTranslation(f.Ru, "PRIVET").With(f.NewPhonetic(f.En, "privet"))).With(
	//		f.NewTranslation(f.Fr, "bonjour")),
	//)
	//
	//deck.With(
	//	f.NewPhrase(
	//		f.NewTranslation(f.En, "goodbye")).With(
	//		f.NewTranslation(f.Ru, "POKA").With(f.NewPhonetic(f.En, "poka"))).With(
	//		f.NewTranslation(f.Fr, "au revouir")),
	//)


	for i:=0; i<10;i++{
		phrase := deck.Next()
		fmt.Println(phrase)
		logrus.Info(phrase.Get(f.Ru))
		logrus.Errorln(deck.Answer(phrase, f.En, "hello"))
	}

}
