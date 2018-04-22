package main

import (
	"flag"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/johnmcdnl/flashcards/routes"
)

//var routes = flag.Bool("routes", false, "Generate router documentation")

func main() {
	flag.Parse()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/deck", func(r chi.Router) {
		r.With(paginate).Get("/", routes.ListPhrases)
		r.Get("/phrases", routes.GetRandomPhrase)
		r.Route("/phrases/{phraseID}", func(r chi.Router) {
			r.Get("/", routes.GetPhrase)
			r.Post("/answer", routes.AttemptAnswer)

			r.Route("/translations", func(r chi.Router) {
				r.Get("/", routes.ListTranslations)
				r.Get("/language/{langCode}", routes.GetTranslationByLang)
			})
		})
	})
	
	http.ListenAndServe(":5800", r)
}

func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// just a stub.. some ideas are to look at URL query params for something like
		// the page number, or the limit, and send a query cursor down the chain
		next.ServeHTTP(w, r)
	})
}
