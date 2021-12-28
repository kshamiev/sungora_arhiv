package general

import (
	"net/http"

	"sungora/lib/logger"
	"sungora/lib/response"
	"sungora/src/config"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

// Ping ping
// @Tags General
// @Summary ping
// @Success 200 {string} string "OK"
// @Router /api/sun/general/ping [get]
func (hh *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)
	rw.JSON("OK")
}

// Version получение версии приложения
// @Tags General
// @Summary получение версии приложения
// @Success 200 {string} string "version"
// @Router /api/sun/general/version [get]
func (hh *Handler) Version(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)
	rw.JSON(config.Get().App.Version)
}

// UploadFile загрузка файла на сервер
// @Tags General
// @Summary загрузка файла на сервер
// @Param file formData file true "загружаемый файл"
// @Accept mpfd
// @Produce octet-stream
// @Success 200 {string} string "OK"
// @Router /api/sun/general/file/upload [post]
func (hh *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)

	data, dataName, err := rw.UploadBuffer()
	if err != nil {
		rw.JSONError(err)
		return
	}
	lg := logger.Gist(r.Context())
	lg.Info(dataName[0])

	rw.Bytes(data[dataName[0]].Bytes(), dataName[0])
}
