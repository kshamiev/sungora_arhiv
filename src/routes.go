package src

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"sungora/lib/logger"
	"sungora/lib/observability"
	"sungora/lib/request"
	"sungora/src/config"
	"sungora/src/handler"
)

// инициализация маршрутов
func initRoutes(cfg *config.App, mux http.Handler) *chi.Mux {
	mid := request.NewMid(cfg.Token, cfg.SigningKey, cfg.DirStatic)

	router := chi.NewRouter()
	router.Use(mid.Cors().Handler)
	router.Use(middleware.Recoverer)
	router.Use(logger.Middleware(logger.Get()))
	router.Use(observability.MiddlewareChi())

	// static
	router.Handle("/template/*", http.FileServer(http.Dir(cfg.DirWork)))

	// swagger
	router.Get("/api/v1/swag/*", mid.Swagger("swagger.json"))

	// grpc gateway
	router.Mount("/", mux)

	// rest
	general(router)

	return router
}

func general(router *chi.Mux) {
	contra := handler.NewGeneral()
	router.Route("/api/v1/general", func(router chi.Router) {
		router.Get("/ping", contra.Ping)
		router.Get("/version", contra.GetVersion)
		router.Get("/test/{id}", contra.Test)
	})
	// websocket
	router.HandleFunc("/api/v1/websocket/gorilla/{id}", contra.GetWebSocketSample)
}
