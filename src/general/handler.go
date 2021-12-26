package general

import (
	"net/http"

	"sungora/lib/response"
	"sungora/src/config"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

// Ping ping
// @Summary ping
// @Tags General
// @Router /api/sun/general/ping [get]
// @Success 200 {string} string "OK"
func (hh *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)
	rw.JSON("OK")
}

// Version получение версии приложения
// @Summary получение версии приложения
// @Tags General
// @Router /api/sun/general/version [get]
// @Success 200 {string} string "version"
func (hh *Handler) Version(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)
	rw.JSON(config.Get().App.Version)
}
