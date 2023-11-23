package server

import (
	"net/http"

	"github.com/BKR-dev/AtollHex/view"
	"github.com/a-h/templ"
)

func timeHandler(componentFunc func() templ.Component) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		component := componentFunc()
		component.Render(r.Context(), w)
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

func NewServer() *http.Server {
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
