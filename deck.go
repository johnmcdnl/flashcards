package flashcards

import "github.com/sirupsen/logrus"

type Deck struct {
	Phrases []*Phrase `json:"phrases"`
}

func NewDeck(p *Phrase) *Deck {
	logrus.Debugln(`func NewDeck(p *Phrase) *Deck {`, p)
	var d Deck

	return d.With(p)
}

func (d *Deck) With(p *Phrase) *Deck {
	logrus.Debugln(`func (d *Deck) With(p *Phrase) *Deck {`, p)
	if p == nil {
		return d
	}
	d.Phrases = append(d.Phrases, p)

	return d
}

func (d *Deck) String() string {
	return jsonString(d)
}
