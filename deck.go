package flashcards

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/boltdb/bolt"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const phrasesBucket = "phrasesBucket"

type Deck struct {
	dbPath   string
	ID       ID       `json:"-"`
	Last     *Card    `json:"-"`
	Current  *Card    `json:"-"`
	Know     Language `json:"know,omitempty"`
	Learning Language `json:"learning,omitempty"`
	Cards    []*Card  `json:"cards,omitempty"`
}

func NewDeck(path string) *Deck {
	if path == "" {
		path = "deck.db"
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
	return &deck
}

func (d *Deck) SaveState() {
	d.saveStateToBucket(phrasesBucket)
}

func (d *Deck) saveStateToBucket(bucketName string) {
	if d.dbPath == "" {
		d.dbPath = "deck.db"
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
func (d *Deck) WithCard(c *Card) *Deck {
	d.Cards = append(d.Cards, c)
	return d
}

func (d *Deck) Next() *Card {
	if len(d.Cards) < 1 {
		panic("no cards!!")
	}

	var totalWeight float64

	for _, card := range d.Cards {
		card.Phrase.Language(d.Learning).Stats.updateWeighting()
		totalWeight += card.Phrase.Language(d.Learning).Stats.Weighting
	}
	r := rand.Intn(int(totalWeight))
	for _, card := range d.Cards {
		r -= int(card.Phrase.Language(d.Learning).Stats.Weighting)
		if r <= 0 {
			d.Last = d.Current
			d.Current = card
			break
		}
	}

	return d.Current
}

func (d *Deck) Shuffle() {

	ShuffleCards(d.Cards)
}

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
