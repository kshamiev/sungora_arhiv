package user

import (
	"sungora/lib/storage/stpg"
	"sungora/lib/worker"

	"github.com/go-chi/chi"
)

func InitDomain(router chi.Router) {
	hh := NewHandler(stpg.Gist())

	// API
	router.Get("/api/sun/users", hh.GetSlice)
	router.Route("/api/sun/user/{id}", func(router chi.Router) {
		router.Post("/", hh.Post)
		router.Put("/", hh.Put)
		router.Get("/", hh.Get)
		router.Delete("/", hh.Delete)
	})
	router.Get("/api/sun/user-test/{id}", hh.Test)

	// HTML

	// CONSOLE
	worker.AddStart(NewTaskOnlineOff(stpg.Gist()))
}
