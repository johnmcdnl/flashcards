package flashcards

type Phrase struct {
	Translations []*Translation
}

func NewPhrase(t *Translation) *Phrase {
	var p Phrase

	return p.With(t)
}

func (p *Phrase) With(t *Translation) *Phrase {
	if t == nil{
		return p
	}
	p.Translations = append(p.Translations, t)
	return p
}
