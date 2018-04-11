package flashcards

import (
	"github.com/sirupsen/logrus"
	"strings"
	"fmt"
)

//Phrase represents a word or sentence that has meaning
type Phrase struct {
	ID                          `json:"id"`
	Translations []*Translation `json:"translations"`
}

// NewPhrase creates a new phrase
func NewPhrase(t *Translation) *Phrase {
	logrus.Debugln(`func NewPhrase(t *Translation) *Phrase {`, t)
	var p Phrase
	p.ID = newID()
	return p.With(t)
}

// With adds Translation to a Phrase
func (p *Phrase) With(t *Translation) *Phrase {
	logrus.Debugln(`func (p *Phrase) With(t *Translation) *Phrase {`, t)
	if t == nil {
		return p
	}

	if exist := p.Get(t.Language); exist !=nil{
		logrus.Error("Phrase with translation is already existing")
		return p
	}

	p.Translations = append(p.Translations, t)
	return p
}

// Get retrieves a Translation for a specific Language
func (p *Phrase) Get(l Language) *Translation {
	for _, t := range p.Translations {
		if t.Language == l {
			return t
		}
	}
	return nil
}

// Answer checks whether an attempt to answer is correct
func (p *Phrase) Answer(l Language, ans string) error {
	logrus.Debugln(`func (p *Phrase) Answer(ans string) error {`, ans)

	correct := p.Get(l).Value

	if !strings.EqualFold(correct, ans) {
		return fmt.Errorf("incorrect! -- expected %s but got %s", correct, ans)
	}

	return nil
}

// String represents a Phrase
func (p *Phrase) String() string {
	return jsonString(p)
}
