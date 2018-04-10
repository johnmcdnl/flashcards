package flashcards

import "strings"

// Phrase is a word or group of words that have a linguistic meaning
// It contains how the linguistic meaning is represented in various languages
type Phrase struct {
	Translations []*Translation
	Stats        []*Stats
}

//NewPhrase returns a new phrase
func NewPhrase(t *Translation) *Phrase {
	return new(Phrase).WithTranslation(t)
}

//WithTranslation adds a version of the phrase with a specific language
func (p *Phrase) WithTranslation(t *Translation) *Phrase {
	if t == nil {
		return p
	}
	if existing := p.GetTranslation(t.Language); existing != nil {
		existing.Value = t.Value
		return p
	}

	p.Translations = append(p.Translations, t)
	return p
}

//GetTranslation gets the value of a phrase in a given language
func (p *Phrase) GetTranslation(l Language) *Translation {
	for _, t := range p.Translations {
		if t.Language == l {
			return t
		}
	}
	return nil
}

//Attempt gives the user a change to see if they understand a phrase
func (p *Phrase) Attempt(attemptLang, targetLang Language, attempt string) bool {
	isSuccessful := strings.EqualFold(p.GetTranslation(attemptLang).Value, attempt)
	p.GetStats(attemptLang, targetLang).Record(isSuccessful)
	return isSuccessful
}

//GetStats gets a user stats for a src and target language combo
func (p *Phrase) GetStats(src, target Language) *Stats {
	for _, s := range p.Stats {
		if s.Source == src && s.Target == target {
			return s
		}
	}
	s := &Stats{
		Source: src,
		Target: target,
	}
	p.Stats = append(p.Stats, s)
	return s
}
