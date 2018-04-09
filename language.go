package flashcards

// Defined  ISO 639 Languages
var (
	Eng = Language{name: "English", native: "english", isoCode: "en"}
	Rus = Language{name: "Russian", native: "русский", isoCode: "ru"}
	Fra = Language{name: "French", native: "français", isoCode: "fr"}
)

//Language represents a ISO 639 Language
type Language struct {
	name    string
	native  string
	isoCode string
}
