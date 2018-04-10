package flashcards

type Language struct {
	name string
	Code string
}

func (l *Language) String() string {
	return jsonString(l)
}

var (
	En = Language{name: "english", Code: "en"}
	Fr = Language{name: "french", Code: "fr"}
)
