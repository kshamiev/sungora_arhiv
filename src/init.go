package src

import (
	"net/http"
	"net/http/pprof"

	"sungora/lib/app"
	"sungora/lib/request"
	"sungora/lib/storage/pgsql"
	"sungora/lib/worker"
	"sungora/src/chat"
	"sungora/src/general"
	"sungora/src/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Init инициализация приложения
func Init(cfg *app.Config) *chi.Mux {
	mid := request.NewMid(cfg.Token, cfg.SigningKey, cfg.DirStatic)

	router := chi.NewRouter()
	router.Use(mid.Cors().Handler)
	router.Use(middleware.Recoverer)
	router.Use(mid.Logger())
	router.Use(mid.Observation())
	router.NotFound(mid.Static(cfg.DirWww))

	// business
	initChat(router)
	initGeneral(router)
	initUser(router)

	// swagger
	router.Get("/api/sun/swag/*", httpSwagger.Handler())

	// static
	router.Handle("/gorilla/*", http.FileServer(http.Dir(cfg.DirWww)))
	router.Handle("/assets/*", http.FileServer(http.Dir(cfg.DirWww)))

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

func initChat(router *chi.Mux) {
	hh := chat.NewHandler()
	// API
	router.HandleFunc("/api/sun/websocket/gorilla/{id}", hh.WebSocketSample)
}

func initGeneral(router *chi.Mux) {
	hh := general.NewHandler()
	// generate html
	router.Get("/", hh.PageIndex)
	router.Get("/index.html", hh.PageIndex)
	router.Get("/page/*", hh.PagePage)
	// API
	router.Route("/api/sun/general", func(router chi.Router) {
		router.Get("/ping", hh.Ping)
		router.Get("/version", hh.Version)
		router.Post("/file/upload", hh.UploadFile)
	})
}

func initUser(router *chi.Mux) {
	hh := user.NewHandler()
	// API
	router.Get("/api/sun/users", hh.GetSlice)
	router.Route("/api/sun/user/{id}", func(router chi.Router) {
		router.Post("/", hh.Post)
		router.Put("/", hh.Put)
		router.Get("/", hh.Get)
		router.Delete("/", hh.Delete)
	})
	router.Get("/api/sun/user-test/{id}", hh.Test)
	// Worker
	worker.AddStart(user.NewTaskOnlineOff(pgsql.Gist()))
}
