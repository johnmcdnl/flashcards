package flashcards

type Stats struct {
	Attempts   int     `json:"attempts"`
	Correct    int     `json:"correct"`
	Percentage float64 `json:"percentage"`
}

func NewStats() *Stats {
	return &Stats{}
}

func (s *Stats) CorrectAttempt() {
	s.Attempts++
	s.Correct++
	s.update()
}

func (s *Stats) WrongAttempt() {
	s.Attempts++
	s.update()
}

func (s *Stats) update() {
	s.updatePercentage()
}

func (s *Stats) updatePercentage() {
	if s.Attempts == 0 {
		s.Percentage = 0
	}
	if s.Correct == 0 {
		s.Percentage = 0
	}
	s.Percentage = float64(s.Correct) / float64(s.Attempts)
}