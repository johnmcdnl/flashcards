package flashcards

// Phrase is a word or group of words that have a linguistic meaning
// It contains how the linguistic meaning is represented in various languages
type Phrase struct {
	Translations []*Translation
}

//WithTranslation adds a version of the phrase with a specific language
func (p *Phrase) WithTranslation(t *Translation) *Phrase {
	if existing := p.GetTranslation(t.Language); existing != nil {
		existing.Value = t.Value
		return p
	}

	p.Translations = append(p.Translations, t)
	return p
}

func (p *Phrase) GetTranslation(l Language) *Translation {
	for _, t := range p.Translations {
		if t.Language == l {
			return t
		}
	}
	return nil
}
