package flashcards

import "math/rand"

type Deck struct {
	last  *Card
	Cards []*Card `json:"cards,omitempty"`
}

func NewDeck() *Deck {
	return &Deck{}
}

func (d *Deck) WithCard(c *Card) *Deck {
	d.Cards = append(d.Cards, c)
	return d
}

func (d *Deck) Next() *Card {
	d.last = d.Cards[0]
	d.Shuffle()
	if len(d.Cards) > 1 && d.Cards[0] == d.last {
		return d.Next()
	}
	return d.Cards[0]
}

func (d *Deck) Shuffle() {
	for i := range d.Cards {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}
