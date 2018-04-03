package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
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

	rand.Seed(42)
	toJson(deck.Next().Phrase.Language(Russian))
	toJson(deck.Next().Phrase.Language(Russian))
	toJson(deck.Next().Phrase.Language(Russian))
	toJson(deck.Next().Phrase.Language(Russian))
	toJson(deck.Next().Phrase.Language(Russian))
	toJson(deck.Next().Phrase.Language(Russian))
	toJson(deck.Next().Phrase.Language(Russian))
	toJson(deck.Next().Phrase.Language(Russian))
	toJson(deck.Next().Phrase.Language(Russian))
	toJson(deck.Next().Phrase.Language(Russian))
}

func toJson(i interface{}) {
	j, _ := json.Marshal(i)
	ji, _ := json.MarshalIndent(i, "\t", " ")
	logrus.Infoln(string(j))
	ioutil.WriteFile("data.json", ji, os.ModePerm)
}
