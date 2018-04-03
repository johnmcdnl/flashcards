package flashcards

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Deck struct {
	Last     *Card    `json:"-"`
	Current  *Card    `json:"-"`
	Know     Language `json:"know,omitempty"`
	Learning Language `json:"learning,omitempty"`
	Cards    []*Card  `json:"cards,omitempty"`
}

func NewDeck() *Deck {

	return &Deck{}
}

func (d *Deck) WithCard(c *Card) *Deck {
	d.Cards = append(d.Cards, c)
	return d
}

func (d *Deck) Next() *Card {
	d.Last = d.Cards[0]
	d.Shuffle()
	d.Current = d.Cards[0]
	if len(d.Cards) > 1 && d.Current == d.Last {
		d.Current = d.Cards[1]
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
