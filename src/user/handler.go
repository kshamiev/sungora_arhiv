package user

import (
	"net/http"

	"sungora/lib/logger"
	"sungora/lib/response"
	"sungora/lib/storage/pgsql"
	"sungora/lib/typ"
	"sungora/src/client"

	"github.com/go-chi/chi"
	"github.com/golang/protobuf/ptypes/empty"
)

type Handler struct {
}

func NewHandler(router *chi.Mux) *Handler {
	hh := &Handler{}
	router.Get("/api/sun/users", hh.Test)
	router.Route("/api/sun/user/{id}", func(router chi.Router) {
		router.Post("/", hh.Post)
		router.Put("/", hh.Put)
		router.Get("/", hh.Get)
		router.Delete("/", hh.Delete)
	})
	router.Get("/api/sun/user-test/{id}", hh.Test)
	return hh
}

func (hh *Handler) GetSlice(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)
	rw.JSON("OK")
}

func (hh *Handler) Get(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)
	rw.JSON("OK")
}

func (hh *Handler) Post(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)
	rw.JSON("OK")
}

func (hh *Handler) Put(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)
	rw.JSON("OK")
}

func (hh *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)
	rw.JSON("OK")
}

// Test
// @Summary test user
// @Tags Users
// @Router /api/sun/user-test/{id} [get]
// @Param id path string true "ID"
// @Success 200 {object} mdsun.User "user"
func (hh *Handler) Test(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)

	usM := NewModel(pgsql.Gist())
	us, err := usM.Load(r.Context(), typ.UUIDMustParse(chi.URLParam(r, "id")))
	if err != nil {
		rw.JSONError(err)
		return
	}

	lg := logger.GetLogger(r.Context())
	lg.Info("General.Test")

	cli := client.GistSunGRPC()
	if _, err := cli.Ping(r.Context(), &empty.Empty{}); err != nil {
		rw.JSONError(err)
		return
	}

	rw.JSON(us)
}
