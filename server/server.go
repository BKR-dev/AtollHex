package server

import (
	"net/http"

	"github.com/BKR-dev/AtollHex/view"

	"github.com/a-h/templ"
	"github.com/go-chi/chi"
)

func NewServer() *http.Server {
	component := view.PageView("Homepage")
	timeComponent := view.ReturnTime()
	articleComponent := view.GetArticles()

	r := chi.NewRouter()
	// serves the hello component at root
	r.Get("/", templ.Handler(component).ServeHTTP)
	// serves only the timestamp div
	r.Get("/time", templ.Handler(timeComponent).ServeHTTP)
	r.Get("/article", templ.Handler(articleComponent).ServeHTTP)

	// create a new http.Server
	server := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	return server
}
