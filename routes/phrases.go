package routes

import (
	"github.com/johnmcdnl/flashcards"
	"net/http"
	"github.com/go-chi/render"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

var deck *flashcards.Deck

func init() {
	deck = flashcards.NewDeck(nil)
}

func ListPhrases(w http.ResponseWriter, r *http.Request) {
	if err := render.RenderList(w, r, NewPhraseListResponse(deck.Phrases)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func GetRandomPhrase(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, NewPhraseResponse(deck.Next()))
}

func GetPhrase(w http.ResponseWriter, r *http.Request) {

	p := deck.GetPhraseByID(chi.URLParam(r, "phraseID"))

	if err := render.Render(w, r, NewPhraseResponse(p)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func AttemptAnswer(w http.ResponseWriter, r *http.Request) {

	p := &flashcards.Phrase{}
	l := flashcards.En
	ans := r.URL.Query().Get("answer")

	err := deck.Answer(p, l, ans)
	logrus.Println(err)
	render.Render(w, r, NewPhraseResponse(p))
}

type PhraseResponse struct {
	*flashcards.Phrase
}

func NewPhraseResponse(phrase *flashcards.Phrase) *PhraseResponse {
	resp := &PhraseResponse{Phrase: phrase}

	return resp
}

func (rd *PhraseResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewPhraseListResponse(phrases []*flashcards.Phrase) []render.Renderer {
	var list []render.Renderer
	for _, phrase := range phrases {
		list = append(list, NewPhraseResponse(phrase))
	}
	return list
}
