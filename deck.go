package flashcards

import "math/rand"

type Deck struct {
	Last    *Card
	Current *Card
	Cards   []*Card `json:"cards,omitempty"`
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
	if len(d.Cards) > 1 && d.Cards[0] == d.Last {
		return d.Next()
	}
	d.Current = d.Cards[0]
	return d.Current
}

func (d *Deck) Shuffle() {
	for i := range d.Cards {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}
