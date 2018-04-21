package main

import (
	"flag"
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/docgen"
	"github.com/go-chi/render"
	"github.com/johnmcdnl/flashcards"
)

var routes = flag.Bool("routes", false, "Generate router documentation")

var deck *flashcards.Deck

func init() {
	deck = flashcards.NewDeck(nil)
}

func main() {
	flag.Parse()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/deck", func(r chi.Router) {
		r.With(paginate).Get("/", ListPhrases)

		r.Post("/phrase", CreatePhrase)
		r.Get("/phrase", GetRandomPhrase)

		r.Route("/phrase/{id}", func(r chi.Router) {
			r.Get("/", GetPhrase)
			r.Put("/", UpdatePhrase)
			r.Delete("/", DeletePhrase)

			r.Post("/answer", AttemptAnswer)
		})
	})

	if *routes {
		fmt.Println(docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
			ProjectPath: "github.com/go-chi/chi",
			Intro:       "Welcome to the chi/_examples/rest generated docs.",
		}))
		return
	}

	http.ListenAndServe(":3333", r)
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

func CreatePhrase(w http.ResponseWriter, r *http.Request) {

	phrase := &flashcards.Phrase{}
	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewPhraseResponse(phrase))
}

func GetPhrase(w http.ResponseWriter, r *http.Request) {
	//TODO
	p := deck.Phrases[0]

	if err := render.Render(w, r, NewPhraseResponse(p)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func UpdatePhrase(w http.ResponseWriter, r *http.Request) {
	//TODO
	p := deck.Phrases[0]
	render.Render(w, r, NewPhraseResponse(p))
}

func DeletePhrase(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func AttemptAnswer(w http.ResponseWriter, r *http.Request) {

	p := &flashcards.Phrase{}
	l := flashcards.En
	ans := "something"
	deck.Answer(p, l, ans)
	render.Render(w, r, NewPhraseResponse(p))
}

func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// just a stub.. some ideas are to look at URL query params for something like
		// the page number, or the limit, and send a query cursor down the chain
		next.ServeHTTP(w, r)
	})
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

type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText string `json:"status"`
	AppCode    int64  `json:"code,omitempty"`
	ErrorText  string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}
