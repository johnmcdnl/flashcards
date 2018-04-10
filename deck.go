package flashcards

type Deck struct {
	//ID                `json:"id"`
	Phrases []*Phrase `json:"phrases"`
}

func NewDeck(p *Phrase) *Deck {
	var d Deck
	//d.ID = newID()

	return d.With(p)
}

func (d *Deck) String() string {
	return jsonString(d)
}

func (d *Deck) With(p *Phrase) *Deck {
	if p == nil {
		return d
	}
	d.Phrases = append(d.Phrases, p)

	return d
}
