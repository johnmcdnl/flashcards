package flashcards

type Stats struct {
	ID         ID      `json:"-"`
	Attempts   int     `json:"attempts"`
	Correct    int     `json:"correct"`
	Percentage float64 `json:"percentage"`
	Weighting  float64 `json:"weighting"`
}

func NewStats() *Stats {
	return &Stats{ID: NewID(), Weighting: 50}
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
	s.updateWeighting()
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

func (s *Stats) updateWeighting() {

	s.updatePercentage()
	s.Weighting = (1 - s.Percentage) * 100

	if s.Attempts <= 5 {
		s.Weighting = 50
	}

	if s.Weighting == 0 {
		s.Weighting = 1
	}
}
