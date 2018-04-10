package flashcards

type Stats struct {
	Source   Language
	Target   Language
	Attempts int
	Correct  int
}

func (s *Stats) Record(b bool) {
	s.Attempts++
	if b {
		s.Correct++
	}
}
