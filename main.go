package main

import (
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

func main() {
	component := hello("Valentine")
	timeComponent := returnTime()

	r := chi.NewRouter()
	// serves the hello component at root
	r.Get("/", templ.Handler(component).ServeHTTP)
	// serves only the timestamp div
	r.Get("/time", templ.Handler(timeComponent).ServeHTTP)

	// starting the server
	println("Listening on :3000")
	http.ListenAndServe(":3000", r)
}

func GetCurrentTime() time.Time {
	return time.Now().UTC()
}
