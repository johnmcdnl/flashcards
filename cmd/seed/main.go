package main

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	. "github.com/johnmcdnl/flashcards"
	"github.com/sirupsen/logrus"
)

func main() {
	parseVerbs()
	parseNouns()
	parseGeneralServiceList()
}

func parseAll(dir string, source, target Language) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".csv") {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				logrus.Fatal(err)
			}
			parse(source, target, data)
		}
		return nil
	})
}

func parseVerbs() {

	data, err := ioutil.ReadFile("./cmd/seed/verbs.csv")
	if err != nil {
		logrus.Fatal(err)
	}
	parse(English, Russian, data)

}

func parseNouns() {

	data, err := ioutil.ReadFile("./cmd/seed/nouns.csv")
	if err != nil {
		logrus.Fatal(err)
	}
	parse(English, Russian, data)

}

func parseGeneralServiceList() {

	data, err := ioutil.ReadFile("./cmd/seed/GeneralServiceList.csv")
	if err != nil {
		logrus.Fatal(err)
	}
	parse(English, Russian, data)

}

func parse(source, target Language, data []byte) {

	records, err := csv.NewReader(bytes.NewReader(data)).ReadAll()
	if err != nil {
		logrus.Fatal(err)
	}

	var deck = NewDeck("deck.db")
	logrus.Infof("Len of deck: %d", len(deck.Cards))

	for _, row := range records {

		var needsToSeed bool = true

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
			NewCard(NewPhrase().WithTranslation(
				NewTranslation(source, row[0])).WithTranslation(
				NewTranslation(target, row[1]).WithPhonetic(
					NewPhonetic(source, row[2])))))

	}
	deck.SaveState()
}
