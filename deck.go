package flashcards

import (
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
	"path/filepath"
	"os"
	"strings"
	"io/ioutil"
	"encoding/csv"
	"bytes"
	"fmt"
	"encoding/json"
)

type Deck struct {
	Phrases []*Phrase `json:"phrases"`
}

func NewDeck(p *Phrase) *Deck {
	logrus.Debugln(`func NewDeck(p *Phrase) *Deck {`, p)
	var d Deck
	d.SeedFromCSV()

	j, _ := json.Marshal(d)
	ioutil.WriteFile("sample-deck.json", j, os.ModePerm)

	return d.With(p)
}

func (d *Deck) With(p *Phrase) *Deck {
	logrus.Debugln(`func (d *Deck) With(p *Phrase) *Deck {`, p)
	if p == nil {
		return d
	}
	d.Phrases = append(d.Phrases, p)

	return d
}

func (d *Deck) Next() *Phrase {
	logrus.Debugln(`func (d *Deck) Next() *Phrase {`, )

	rand.Seed(time.Now().UnixNano())
	return d.Phrases[rand.Intn(len(d.Phrases))]
}

func (d *Deck) Answer(p *Phrase, l Language, ans string) error {
	logrus.Debugln(`func (d *Deck) Answer(p *Phrase, answer string) error {`, p, l, ans)

	return p.Answer(l, ans)
}

func (d *Deck) SeedFromCSV() {
	walkFn := func(path string, _ os.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".csv") {
			return nil
		}

		return d.seedFromCSV(path)
	}
	filepath.Walk("./", walkFn)
}

func (d *Deck) seedFromCSV(path string) error {

	records, err := readCSV(path)
	if err != nil {
		return err
	}

	if len(records) <= 2 {
		return fmt.Errorf("not enough data")
	}
	headers := records[0]
	data := records[1:]

	fmt.Println(len(d.Phrases))
	fmt.Println(len(d.Phrases))
	fmt.Println(len(d.Phrases))
	for _, row := range data {
		phrase := NewPhrase(nil)
		for i, col := range row {
			val := strings.Trim(col, " ")

			lang := headers[i]
			if len(strings.Trim(lang, " ")) != 2 {
				//It's a phonetic
				split := strings.Split(lang, "_")
				phraseLang := GetLang(strings.Replace(split[0],"_", "", -1))
				phoneticLang := GetLang(strings.Replace(split[1],"_", "", -1))

				phrase.Get(phraseLang).With(NewPhonetic(phoneticLang, val))
				continue
			}

			phrase.With(NewTranslation(GetLang(lang), val))
		}
		d.With(phrase)
	}

	fmt.Println(len(d.Phrases))
	fmt.Println(len(d.Phrases))
	fmt.Println(len(d.Phrases))

	fmt.Println(headers)
	fmt.Println(data)

	return nil
}

func readCSV(path string) ([][]string, error) {

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(bytes.NewReader(b))

	return reader.ReadAll()
}

func (d *Deck) String() string {
	return jsonString(d)
}
