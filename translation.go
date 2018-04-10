package flashcards

type Translation struct {
	Language
	Value     string
	Phonetics []*Phonetic
}

func NewTranslation(p *Phonetic)*Translation{
	var t Translation

	return t.With(p)
}

func (t *Translation)With(p *Phonetic)*Translation{
	if p == nil{
		return t
	}
	t.Phonetics = append(t.Phonetics, p)
	return t
}