package flashcards

//Translation represents a phrase in language
type Translation struct {
	Language Language
	Value    string
}

//NewTranslation produces a translation for a language
func NewTranslation(l Language, v string) *Translation {
	return &Translation{
		Language: l,
		Value:    v,
	}
}
