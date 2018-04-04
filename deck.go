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
		logrus.Fatal(err)
	}
	defer db.Close()
	var deck Deck
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("deck"))
		if err != nil {
			return err
		}
		deckBytes := bucket.Get([]byte("deck"))
		if len(deckBytes) == 0 {
			deck = Deck{ID: NewID()}
			return nil
		}
		return json.Unmarshal(deckBytes, &deck)
	})
	if err != nil {
		logrus.Fatal(err)
	}
	deck.dbPath = path
	return &deck
}

func (d *Deck) SaveState() {
	if d.dbPath == "" {
		d.dbPath = "deck.db"
	}
	db, err := bolt.Open(d.dbPath, os.ModePerm, &bolt.Options{ReadOnly: false})
	if err != nil {
		logrus.Fatal(err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("deck"))
		if err != nil {
			return err
		}
		data, err := json.Marshal(d)
		if err != nil {
			return err
		}
		return bucket.Put([]byte("deck"), data)
	})
	if err != nil {
		logrus.Fatal(err)
	}
}

func (d *Deck) WithCard(c *Card) *Deck {
	d.Cards = append(d.Cards, c)
	return d
}

func (d *Deck) Next() *Card {

	var totalWeight float64
	for _, card := range d.Cards {
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
