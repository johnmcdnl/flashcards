package flashcards

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"go/build"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/boltdb/bolt"
	"github.com/sirupsen/logrus"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const phrasesBucket = "phrasesBucket"

// DefaultDeckName is the default deck db
const DefaultDeckName = "flashcardsDeck.db"

// Deck is a parennt for a group of cards
type Deck struct {
	dbPath   string
	ID       ID       `json:"-"`
	Last     *Card    `json:"-"`
	Current  *Card    `json:"-"`
	Know     Language `json:"know,omitempty"`
	Learning Language `json:"learning,omitempty"`
	subdeck  []*Card
	Cards    []*Card `json:"cards,omitempty"`
}

// NewDeckWithSize extracts a subset of NewDeck
func NewDeckWithSize(path string, start, end int, shuffle bool) *Deck {
	deck := NewDeck(path)

	if shuffle {
		deck.Shuffle()
	}

	deck.subdeck = deck.Cards[start:end]

	return deck
}

// NewStandardDeck returns a deck with sensible defaults
func NewStandardDeck() *Deck {
	return NewDeck(DefaultDeckName)
}

// NewDeck generates a deck with all known phrases
func NewDeck(path string) *Deck {
	d := newDeck(path)
	if len(d.Cards) == 0 {
		d.seed()
		return NewDeck(path)
	}
	return d
}

func newDeck(path string) *Deck {
	if path == "" {
		path = DefaultDeckName
	}
	db, err := bolt.Open(path, os.ModePerm, &bolt.Options{ReadOnly: false})
	if err != nil {
		logrus.Fatalf("NewDeck %s", err)
	}
	defer db.Close()
	var deck Deck
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(phrasesBucket))
		if err != nil {
			return err
		}
		deckBytes := bucket.Get([]byte(phrasesBucket))
		if len(deckBytes) == 0 {
			deck = Deck{ID: NewID()}
			return nil
		}
		return json.Unmarshal(deckBytes, &deck)
	})
	if err != nil {
		logrus.Fatalf("NewDeck 2 %s", err)
	}
	deck.dbPath = path
	deck.Know = English
	deck.Learning = Russian
	return &deck
}

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

// SaveState saves the current state of the deck
func (d *Deck) SaveState() {
	d.saveStateToBucket(phrasesBucket)
}

func (d *Deck) saveStateToBucket(bucketName string) {
	if d.dbPath == "" {
		d.dbPath = DefaultDeckName
	}

	db, err := bolt.Open(d.dbPath, os.ModePerm, &bolt.Options{ReadOnly: false})
	if err != nil {
		logrus.Fatalf("SaveState %s", err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		data, err := json.Marshal(d)
		if err != nil {
			return err
		}
		return bucket.Put([]byte(bucketName), data)
	})
	if err != nil {
		logrus.Fatalf("SaveState 2 %s", err)
	}
}

// WithCard adds a new card to the deck
func (d *Deck) WithCard(c *Card) *Deck {
	d.Cards = append(d.Cards, c)
	return d
}

// Next takes the first card from the deck
func (d *Deck) Next() *Card {

	if len(d.subdeck) > 0 {
		return d.next(d.subdeck)
	}

	return d.next(d.Cards)
}

func (d *Deck) next(cards []*Card) *Card {
	if len(cards) < 1 {
		panic("no cards!!")
	}

	var totalWeight float64

	for _, card := range cards {
		card.Phrase.Language(d.Learning).Stats.updateWeighting()
		totalWeight += card.Phrase.Language(d.Learning).Stats.Weighting
	}
	r := rand.Intn(int(totalWeight))
	for _, card := range cards {
		r -= int(card.Phrase.Language(d.Learning).Stats.Weighting)
		if r <= 0 {
			d.Last = d.Current
			d.Current = card
			break
		}
	}

	return d.Current
}

// Shuffle reorders all cards in the deck
func (d *Deck) Shuffle() {

	ShuffleCards(d.Cards)
}

// GetIncorrectGuesses gets a set of incorrect answers
func (d *Deck) GetIncorrectGuesses(correct *Card, n int) []*Card {

	if len(d.Cards) == n {
		return d.Cards
	}

	var copyCards []*Card

	for _, c := range d.Cards {
		if c == correct {
			continue
		}
		copyCards = append(copyCards, c)
	}

	ShuffleCards(copyCards)
	return copyCards[:n]
}
