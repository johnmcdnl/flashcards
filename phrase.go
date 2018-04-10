package flashcards

import "github.com/sirupsen/logrus"

type Phrase struct {
	ID                          `json:"id"`
	Translations []*Translation `json:"translations"`
}

func NewPhrase(t *Translation) *Phrase {
	logrus.Debugln(`func NewPhrase(t *Translation) *Phrase {`, t)
	var p Phrase
	p.ID = newID()
	return p.With(t)
}

func (p *Phrase) With(t *Translation) *Phrase {
	logrus.Debugln(`func (p *Phrase) With(t *Translation) *Phrase {`, t)
	if t == nil {
		return p
	}
	p.Translations = append(p.Translations, t)
	return p
}

func (p *Phrase) String() string {
	return jsonString(p)
}
