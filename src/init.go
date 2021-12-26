package src

import (
	"net/http"
	"net/http/pprof"

	"sungora/lib/worker"
	"sungora/src/chat"
	"sungora/src/config"
	"sungora/src/general"
	"sungora/src/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	"sungora/lib/request"
	_ "sungora/template/swagger"
)

// Init инициализация приложения
func Init(cfg *config.App) *chi.Mux {
	mid := request.NewMid(cfg.Token, cfg.SigningKey, cfg.DirStatic)

	router := chi.NewRouter()
	router.Use(mid.Cors().Handler)
	router.Use(middleware.Recoverer)
	router.Use(mid.Logger())
	router.Use(mid.Observation())

	// static
	router.Handle("/template/*", http.FileServer(http.Dir(cfg.DirWork)))

	// swagger
	router.Get("/api/sun/swag/*", httpSwagger.Handler())

	// business

	user.NewHandler(router)
	worker.AddStart(user.NewTaskOnlineOff())

	general.NewHandler(router)

	chat.NewHandler(router)

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
