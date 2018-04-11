package flashcards

import "github.com/sirupsen/logrus"

// Translation represents a phrase in a Language
type Translation struct {
	//ID                    `json:"id"`
	Language
	Value     string      `json:"value"`
	Phonetics []*Phonetic `json:"phonetics,omitempty"`
}

// NewTranslation generates a new Translation
func NewTranslation(l Language, v string) *Translation {
	logrus.Debugln(`func NewTranslation(l Language, v string) *Translation {`, l, v)
	return &Translation{
		//ID: newID(),
		Language: l,
		Value:    v,
	}
}

// With Phonetic adds a phonetic to a Translation
func (t *Translation) With(p *Phonetic) *Translation {
	logrus.Debugln(`func (t *Translation) With(p *Phonetic) *Translation {`, p)
	if p == nil {
		return t
	}

	if exist := t.Get(p.Language); exist != nil {
		logrus.Errorf("p Phonetic looks to be pre-existing")
		return t
	}

	t.Phonetics = append(t.Phonetics, p)
	return t
}

// Get retrieves a specific translation based on Language
func (t *Translation) Get(l Language) *Phonetic {
	for _, p := range t.Phonetics {
		if t.Language == l {
			return p
		}
	}
	return nil
}

// String represents a Translation
func (t *Translation) String() string {
	return jsonString(t)
}
