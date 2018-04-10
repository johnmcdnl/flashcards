package flashcards

type Translation struct {
	Language
	Value     string
	Phonetics []*Phonetic
}
