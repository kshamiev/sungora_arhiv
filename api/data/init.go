package data

import (
	"sungora/lib/storage/stpg"
	"sungora/lib/worker"

	"github.com/go-chi/chi"
)

func InitDomain(router *chi.Mux) {
	hh := NewHandler(stpg.Gist())

	// API
	router.Route("/api/sun/data", func(router chi.Router) {
		router.Post("/upload-test", hh.UploadFile)
		router.Post("/upload", hh.Upload)
	})

	// HTML

	// CONSOLE
	worker.AddStart(NewTaskStorageClear(stpg.Gist()))
}
