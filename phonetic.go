package flashcards

import "github.com/sirupsen/logrus"

type Phonetic struct {
	Language
	Value string `json:"value"`
}

func NewPhonetic(l Language, v string) *Phonetic {
	logrus.Debugln(`func NewPhonetic(l Language, v string) *Phonetic {`, l, v)
	return &Phonetic{
		Language: l,
		Value:    v,
	}
}

func (p *Phonetic) String() string {
	return jsonString(p)
}
