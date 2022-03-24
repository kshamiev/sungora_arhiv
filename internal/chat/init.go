package chat

import "github.com/go-chi/chi"

func InitDomain(router chi.Router) {
	hh := NewHandler()

	// API
	router.HandleFunc("/api/sun/websocket/gorilla/{id}", hh.WebSocketSample)

	// HTML

	// CONSOLE
}
