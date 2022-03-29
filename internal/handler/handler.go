package handler

import (
	"net/http"
	"net/http/pprof"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	"sample/internal/config"
	"sample/internal/task"
	"sample/lib/app"
	"sample/lib/app/request"
	"sample/lib/app/sheduler"
	"sample/lib/jaeger"
	"sample/lib/logger"
	"sample/lib/storage"
	"sample/lib/storage/stpg"
	"sample/lib/tpl"
)

type Handler struct {
	st    storage.Face
	wsBus app.SocketBus
}

func NewHandler() *Handler {
	return &Handler{
		st:    stpg.Gist(),
		wsBus: app.NewWebSocketBus(),
	}
}

func Routing(cfg *config.App) *chi.Mux {
	router := chi.NewRouter()

	if cfg.Mode == "dev" {
		// swagger
		router.Get("/api/sun/swag/*", httpSwagger.Handler())
		// pprof
		router.Get("/api/sun/debug/pprof/index", pprof.Index)
		router.Get("/api/sun/debug/pprof/cmdline", pprof.Cmdline)
		router.Get("/api/sun/debug/pprof/profile", pprof.Profile)
		router.Get("/api/sun/debug/pprof/symbol", pprof.Symbol)
		router.Get("/api/sun/debug/pprof/trace", pprof.Trace)
		router.Get("/api/sun/debug/pprof/allocs", pprof.Handler("allocs").ServeHTTP)
		router.Get("/api/sun/debug/pprof/heap", pprof.Handler("heap").ServeHTTP)
		router.Get("/api/sun/debug/pprof/goroutine", pprof.Handler("goroutine").ServeHTTP)
		router.Get("/api/sun/debug/pprof/threadcreate", pprof.Handler("threadcreate").ServeHTTP)
		router.Get("/api/sun/debug/pprof/block", pprof.Handler("block").ServeHTTP)
		router.Get("/api/sun/debug/pprof/mutex", pprof.Handler("mutex").ServeHTTP)
	}

	// static
	router.Handle("/assets/*", http.FileServer(http.Dir(cfg.DirWww)))

	mid := request.NewMid(cfg.Token, cfg.SigningKey)
	router.Group(func(r chi.Router) {
		r.Use(mid.Cors().Handler)
		r.Use(middleware.Recoverer)
		r.Use(jaeger.Observation())
		r.Use(logger.Middleware())
		hh := NewHandler()

		// user
		r.Get("/sun/api/v1/users", hh.GetSlice)
		r.Route("/sun/api/v1/user/{id}", func(r chi.Router) {
			r.Post("/", hh.Post)
			r.Put("/", hh.Put)
			r.Get("/", hh.Get)
			r.Delete("/", hh.Delete)
		})
		r.Get("/sun/api/v1/user-sample/{id}", hh.Sample)

		// chat
		r.HandleFunc("/sun/api/v1/websocket/gorilla/{id}", hh.WebSocketSample)

		// data
		r.Route("/sun/api/v1/data", func(router chi.Router) {
			router.Post("/upload-test", hh.UploadFile)
			router.Post("/upload", hh.Upload)
			router.Get("/download/{id}", hh.Download)
		})

		// general
		r.Route("/sun/api/v1/general", func(r chi.Router) {
			r.Get("/ping", hh.Ping)
			r.Get("/version", hh.Version)
		})

		// html
		r.Get("/", hh.PageIndex)
		r.Get("/index.html", hh.PageIndex)
		r.Get("/page/*", hh.PagePage)
	})

	// task bg
	sheduler.AddStart(task.NewTaskOnlineOff())
	sheduler.AddStart(task.NewTaskStorageClear())
	sheduler.AddStart(tpl.NewTaskTemplateParse(cfg.DirWww))

	return router
}
