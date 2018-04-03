package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"

	. "github.com/johnmcdnl/flashcards"
)

func main() {
	deck := NewDeck().WithCard(
		NewCard(NewPhrase().WithTranslation(
			NewTranslation(English, "Hello")).WithTranslation(
			NewTranslation(Russian, "Здравствуйте").WithPhonetic(
				NewPhonetic(English, "Zdravstvuyte"))).WithTranslation(
			NewTranslation(French, "Bonjour")),
		)).WithCard(
		NewCard(NewPhrase().WithTranslation(
			NewTranslation(English, "Goodbye")).WithTranslation(
			NewTranslation(Russian, "до свидания").WithPhonetic(
				NewPhonetic(English, "do svidaniya"))).WithTranslation(
			NewTranslation(French, "Au revoir")),
		)).WithCard(
		NewCard(NewPhrase().WithTranslation(
			NewTranslation(English, "Dog")).WithTranslation(
			NewTranslation(Russian, "Собака").WithPhonetic(
				NewPhonetic(English, "Sobaka"))).WithTranslation(
			NewTranslation(French, "Chien")),
		)).WithCard(
		NewCard(NewPhrase().WithTranslation(
			NewTranslation(English, "Elephant")).WithTranslation(
			NewTranslation(Russian, "слон").WithPhonetic(
				NewPhonetic(English, "slon"))).WithTranslation(
			NewTranslation(French, "l'éléphant")),
		),
	)

	deck.Know = English
	deck.Learning = Russian
	for i := 1; i <= 5; i++ {
		log(deck.Next().Phrase.Language(Russian))
		logError(deck.Current)

		deck.Current.PrintQuestion(deck)
		deck.Current.AttemptAnswer(deck.Know, deck.Learning, "")
	}

	toJson(deck)

}

func toJson(i interface{}) {
	j, _ := json.Marshal(i)
	ji, _ := json.MarshalIndent(i, " ", "\t")
	logrus.Debugln(string(j))
	ioutil.WriteFile("data.json", ji, os.ModePerm)
}

func logError(i interface{}) {
	j, _ := json.Marshal(i)
	logrus.Debug(string(j))
	fmt.Println()
}

func log(i interface{}) {
	j, _ := json.Marshal(i)
	logrus.Debug(string(j))
}
