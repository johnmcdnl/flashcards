package flashcards

type Language struct {
	name string
	code string
}

var (
	En = Language{name: "english", code: "en"}
	Fr = Language{name: "french", code: "fr"}
)
