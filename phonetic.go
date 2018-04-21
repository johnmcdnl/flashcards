package flashcards

import "github.com/sirupsen/logrus"

//Phonetic is a pronunciation for a phrase
type Phonetic struct {
	Language
	Value string `json:"value"`
}

//NewPhonetic creates a new phonetic
func NewPhonetic(l Language, v string) *Phonetic {
	logrus.Debugln(`func NewPhonetic(l Language, v string) *Phonetic {`, l, v)
	return &Phonetic{
		Language: l,
		Value:    v,
	}
}

// String represents a Phonetic
func (p *Phonetic) String() string {
	return jsonString(p)
}
