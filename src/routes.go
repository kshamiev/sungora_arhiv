package src

import (
	"context"
	"net/http"
	"net/http/pprof"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"sungora/lib/logger"
	"sungora/lib/observability"
	"sungora/lib/request"
	"sungora/src/config"
	"sungora/src/handler"
)

// инициализация маршрутов
func initRoutes(cfg *config.App) *chi.Mux {
	mid := request.NewMid(cfg.Token, cfg.SigningKey, cfg.DirStatic)

	router := chi.NewRouter()
	router.Use(mid.Cors().Handler)
	router.Use(middleware.Recoverer)
	router.Use(logger.Middleware(logger.Gist(context.Background())))
	router.Use(observability.MiddlewareChi())

	// static
	router.Handle("/template/*", http.FileServer(http.Dir(cfg.DirWork)))

	// swagger
	router.Get("/api/sun/swag/*", mid.Swagger("swagger.json"))

	// rest
	general(router)

	// pprof
	router.Get("/api/sun/debug/pprof/trace", func(w http.ResponseWriter, r *http.Request) {
		pprof.Trace(w, r)
	})
	router.Get("/api/sun/debug/pprof/profile", func(w http.ResponseWriter, r *http.Request) {
		pprof.Profile(w, r)
	})
	router.Get("/api/sun/debug/pprof/symbol", func(w http.ResponseWriter, r *http.Request) {
		pprof.Symbol(w, r)
	})
	router.Get("/api/sun/debug/pprof/allocs", pprof.Handler("allocs").ServeHTTP)
	router.Get("/api/sun/debug/pprof/heap", pprof.Handler("heap").ServeHTTP)
	router.Get("/api/sun/debug/pprof/goroutine", pprof.Handler("goroutine").ServeHTTP)

	return router
}

func general(router *chi.Mux) {
	contra := handler.NewGeneral()
	router.Route("/api/sun/general", func(router chi.Router) {
		router.Get("/ping", contra.Ping)
		router.Get("/version", contra.Version)
		router.Get("/test/{id}", contra.Test)
	})
	// websocket
	router.HandleFunc("/api/sun/websocket/gorilla/{id}", contra.WebSocketSample)
}
