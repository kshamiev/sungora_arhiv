package data

import (
	"sungora/lib/storage/stpg"
	"sungora/lib/worker"

	"github.com/go-chi/chi"
)

func InitDomain(router chi.Router) {
	hh := NewHandler(stpg.Gist())

	// API
	router.Route("/api/sun/data", func(router chi.Router) {
		router.Post("/upload-test", hh.UploadFile)
		router.Post("/upload", hh.Upload)
		router.Get("/download/{id}", hh.Download)
	})

	// HTML

	// CONSOLE
	worker.AddStart(NewTaskStorageClear(stpg.Gist()))
}
