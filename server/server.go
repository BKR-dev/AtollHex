package server

import (
	"net/http"
	"time"

	"github.com/BKR-dev/AtollHex/logic"
	"github.com/BKR-dev/AtollHex/view"
)

// func NewServer() *http.Server {
// 	currentTime := logic.GetCurrentTime().String()
// 	component := view.PageView("AtollHEX")
// 	timeComponent := view.ReturnTime(currentTime)

// 	r := chi.NewRouter()

// 	r.Use(middleware.Logger)
// 	r.Use(middleware.Recoverer)
// 	// serves the hello component at root
// 	r.Get("/", templ.Handler(component).ServeHTTP)
// 	// serves only the timestamp div
// 	r.Get("/time", templ.Handler(timeComponent).ServeHTTP)
// 	// serves articles
// 	r.Get("/articles", func(w http.ResponseWriter, r *http.Request) {
// 		page, _ := strconv.Atoi(r.URL.Query().Get("page"))

// 		// Fetch articles based on page and pageSize
// 		// This is just a placeholder. Replace it with your actual logic.
// 		articles := view.ReturnArticles(page)

// 		// Convert articles to JSON and write to response
// 		// This is also a placeholder. Replace it with your actual logic.
// 		json.NewEncoder(w).Encode(articles)
// 	})

// 	// create a new http.Server
// 	server := &http.Server{
// 		Addr:    ":3000",
// 		Handler: r,
// 	}

// 	return server
// }

// NowHandler struct the returns the current time in a anonymous function
type NowHandler struct {
	Now func() time.Time
}

// NewNowHandler returns a new NowHandler
func NewNowHandler(now func() time.Time) NowHandler {
	return NowHandler{now}
}

// ServeHTTP confirms to the http.Handler interface on the NowHandler struct
func (nh NowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	view.TimeComponent(nh.Now().String()).Render(r.Context(), w)
}

func NewServer() *http.Server {
	currentTime := logic.GetCurrentTime().String()
	pageViewComponent := view.PageView("AtollHEX", currentTime)
	timeComponent := view.ReturnTime(currentTime)
	http.Handle("/", NewNowHandler(time.Now))
}
