package flashcards

import "github.com/sirupsen/logrus"

type Translation struct {
	//ID                    `json:"id"`
	Language
	Value     string      `json:"value"`
	Phonetics []*Phonetic `json:"phonetics,omitempty"`
}

func NewTranslation(l Language, v string) *Translation {
	logrus.Debugln(`func NewTranslation(l Language, v string) *Translation {`, l, v)
	return &Translation{
		//ID: newID(),
		Language: l,
		Value:    v,
	}
}

func (t *Translation) With(p *Phonetic) *Translation {
	logrus.Debugln(`func (t *Translation) With(p *Phonetic) *Translation {`, p)
	if p == nil {
		return t
	}
	t.Phonetics = append(t.Phonetics, p)
	return t
}

func (t *Translation) String() string {
	return jsonString(t)
}
