package flashcards

import (
	"bytes"
	"encoding/csv"
	"go/build"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

var once sync.Once

func (d *Deck) seed() {

	dir := filepath.Join(build.Default.GOPATH,
		"src", "github.com", "johnmcdnl", "flashcards")

	// parseAll(path, English, Russian)

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		logrus.Info("seed.parseAll", dir)
		// panic(dir)
		if strings.HasSuffix(path, ".csv") {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				logrus.Fatal(err)
			}
			// TODO - this should be able to read CSV header column to know lanuage
			d.parse(d.Know, d.Learning, data)
		}
		return nil
	})

}

// TODO - this should be able to read CSV header column to know lanuage
func (d *Deck) parse(source, target Language, data []byte) {

	records, err := csv.NewReader(bytes.NewReader(data)).ReadAll()
	if err != nil {
		logrus.Fatal(err)
	}

	// var deck = NewDeck(DeckName)
	logrus.Infof("Len of deck: %d", len(d.Cards))

	for _, row := range records {

		var needsToSeed = true

		for _, c := range d.Cards {
			if c.Phrase.Language(source).Value == row[0] {
				needsToSeed = false
			}
		}

		if !needsToSeed {
			continue
		}

		logrus.Infoln("Seeding", row)
		d.WithCard(
			NewCard(NewPhrase().WithTranslation(
				NewTranslation(source, row[0])).WithTranslation(
				NewTranslation(target, row[1]).WithPhonetic(
					NewPhonetic(source, row[2])))))

	}
	d.SaveState()
}
