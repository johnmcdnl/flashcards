package flashcards


// Defined Languages
var (
	English = language{name: "English", native: "english", isoCode:"en"}
	Eng = English
	 Russian = language{name: "Russian", native: "русский", isoCode:"ru"}
	 Rus = Russian
	 French = language{name: "French", native: "français", isoCode:"fr"}
	 Fra = French
)

//Language represents a ISO 639 Language
type language struct{
	name string 
	native string 
	isoCode string 
}

