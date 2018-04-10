package flashcards


//Stats represent how well a user is learning a given phrase
type Stats struct {
	Source   Language
	Target   Language
	Attempts int
	Correct  int
}

//Record captures an user's attempt
func (s *Stats) Record(b bool) {
	s.Attempts++
	if b {
		s.Correct++
	}
}
