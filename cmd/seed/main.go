package main

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"

	"github.com/johnmcdnl/flashcards"
	"github.com/sirupsen/logrus"
)

func main() {
	parseVerbs()
	parseNouns()
	parseGeneralServiceList()
}

func parseVerbs() {

	data, err := ioutil.ReadFile("./cmd/seed/verbs.csv")
	if err != nil {
		logrus.Fatal(err)
	}
	parse(flashcards.English, flashcards.Russian, data)

}

func parseNouns() {

	data, err := ioutil.ReadFile("./cmd/seed/nouns.csv")
	if err != nil {
		logrus.Fatal(err)
	}
	parse(flashcards.English, flashcards.Russian, data)

}

func parseGeneralServiceList() {

	data, err := ioutil.ReadFile("./cmd/seed/GeneralServiceList.csv")
	if err != nil {
		logrus.Fatal(err)
	}
	parse(flashcards.English, flashcards.Russian, data)

}

func parse(source, target flashcards.Language, data []byte) {

	records, err := csv.NewReader(bytes.NewReader(data)).ReadAll()
	if err != nil {
		logrus.Fatal(err)
	}

	var deck = flashcards.NewDeck(flashcards.DeckName)
	logrus.Infof("Len of deck: %d", len(deck.Cards))

	for _, row := range records {

		var needsToSeed = true

		for _, c := range deck.Cards {
			if c.Phrase.Language(source).Value == row[0] {
				needsToSeed = false
			}
		}

		if !needsToSeed {
			continue
		}

		logrus.Infoln("Seeding", row)
		deck.WithCard(
			flashcards.NewCard(flashcards.NewPhrase().WithTranslation(
				flashcards.NewTranslation(source, row[0])).WithTranslation(
				flashcards.NewTranslation(target, row[1]).WithPhonetic(
					flashcards.NewPhonetic(source, row[2])))))

	}
	deck.SaveState()
}
