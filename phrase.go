package flashcards

import (
	"github.com/sirupsen/logrus"
	"strings"
	"fmt"
)

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

func (p *Phrase) Get(l Language) *Translation {
	for _, t := range p.Translations {
		if t.Language == l {
			return t
		}
	}
	return nil
}

func (p *Phrase) Answer(l Language, ans string) error {
	logrus.Debugln(`func (p *Phrase) Answer(ans string) error {`, ans)

	correct := p.Get(l).Value

	if !strings.EqualFold(correct, ans) {
		return fmt.Errorf("incorrect! -- expected %s but got %s", correct, ans)
	}

	return nil
}

func (p *Phrase) String() string {
	return jsonString(p)
}
