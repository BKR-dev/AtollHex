package server

import (
	"net/http"

	"github.com/BKR-dev/AtollHex/view"
	"github.com/a-h/templ"
)

// NowHandler struct the returns the current time in a anonymous function
// type NowHandler struct {
// 	Now func() time.Time
// }

// NewNowHandler returns a new NowHandler
// func NewNowHandler(now func() time.Time) NowHandler {
// 	return NowHandler{now}
// }

// ServeHTTP confirms to the http.Handler interface on the NowHandler struct
// func (nh NowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	view.TimeComponent(nh.Now().String()).Render(r.Context(), w)
// }

func timeHandler(componentFunc func() templ.Component) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		component := componentFunc()
		component.Render(r.Context(), w)
	}
}

func NewServer() *http.Server {
	// currentTime := logic.GetCurrentTime().String()
	pageViewComponentFunc := func() templ.Component {
		return view.PageView("AtollHEX")
	}
	pageViewHandler := NewTemplHandler(pageViewComponentFunc)

	http.Handle("/", pageViewHandler)
	http.Handle("/time", http.HandlerFunc(timeHandler(view.ReturnTime)))

	return &http.Server{
		Addr:    ":3000",
		Handler: http.DefaultServeMux,
	}
}

type TemplHandler struct {
	ComponentFunc func() templ.Component
}

func NewTemplHandler(componentFunc func() templ.Component) TemplHandler {
	return TemplHandler{componentFunc}
}

func (th TemplHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	component := th.ComponentFunc()
	component.Render(r.Context(), w)
}

// The one thing I would do is change the struct.
// Keep the New function and keep the interface method,
// but the type that’s a function is going to throw a lot of people off imo
