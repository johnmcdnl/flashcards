package flashcards

type Deck struct {
	Cards []*Card `json:"cards,omitempty"`
}

func NewDeck() *Deck {
	return &Deck{}
}

func (d *Deck) WithCard(c *Card) *Deck {
	d.Cards = append(d.Cards, c)
	return d
}
