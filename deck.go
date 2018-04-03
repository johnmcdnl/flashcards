package flashcards

type Deck struct {
	Cards []*Card `json:"cards,omitempty"`
}

func NewDeck() *Deck {
	return &Deck{}
}
