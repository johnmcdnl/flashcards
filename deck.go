package flashcards

import (
	"encoding/json"
	"fmt"
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
	ID       ID       `json:"id"`
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
	d.Last = d.Current
	d.Shuffle()
	d.Current = d.Cards[0]
	if len(d.Cards) > 1 && d.Current == d.Last {
		d.Current = d.Cards[1]
	}

	return d.Current
}

func (d *Deck) NextWeighted() *Card {
	var stats []*Stats
	for _, c := range d.Cards {
		for _, t := range c.Phrase.Translations {

			if t.Language == d.Learning {
				stats = append(stats, t.Stats)
			}
		}
	}

	// toJSON(stats)
	return d.Next()
}

func toJSON(i interface{}) {
	j, _ := json.Marshal(i)
	fmt.Println(string(j))
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
