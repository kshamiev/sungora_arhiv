package general

import "github.com/go-chi/chi"

func InitDomain(router chi.Router) {
	hh := NewHandler()

	// API
	router.Route("/api/sun/general", func(router chi.Router) {
		router.Get("/ping", hh.Ping)
		router.Get("/version", hh.Version)
	})

	// HTML
	router.Get("/", hh.PageIndex)
	router.Get("/index.html", hh.PageIndex)
	router.Get("/page/*", hh.PagePage)

	// CONSOLE
}
