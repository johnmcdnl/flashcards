package flashcards

type Deck struct {
	Phrases []*Phrase
}

func NewDeck(p *Phrase) *Deck {
	var d Deck

	return d.With(p)
}

func (d *Deck) With(p *Phrase) *Deck {
	if p == nil {
		return d
	}
	d.Phrases = append(d.Phrases, p)

	return d
}
