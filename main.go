package main

import (
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

func main() {
	component := hello("John", getCurrentTime().String())
	timeComponent := returnTime(getCurrentTime().String())

	r := chi.NewRouter()
	// serves the hello component at root
	r.Get("/", templ.Handler(component).ServeHTTP)
	// serves only the timestamp div
	r.Get("/time", templ.Handler(timeComponent).ServeHTTP)
	// serves only the timestamp div
	r.Get("/time2", timeHandler)

	// starting the server
	println("Listening on :3000")
	http.ListenAndServe(":3000", r)
}

func getCurrentTime() time.Time {
	return time.Now().UTC()
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	message := `<div id="timestamp">at current time: `
	message = message + getCurrentTime().String() + `</div>`
	w.Write([]byte(message))
}
