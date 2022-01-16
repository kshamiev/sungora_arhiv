package data

import (
	"net/http"

	"sungora/lib/errs"
	"sungora/lib/logger"
	"sungora/lib/response"
	"sungora/lib/storage"
	"sungora/services/mdsungora"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Handler struct {
	st storage.Face
}

func NewHandler(st storage.Face) *Handler {
	return &Handler{
		st: st,
	}
}

// UploadFile загрузка файла на сервер
// @Tags Data
// @Summary загрузка файла на сервер
// @Param file formData file true "загружаемый файл"
// @Accept mpfd
// @Produce octet-stream
// @Success 200 {string} string "OK"
// @Router /api/sun/data/upload-test [post]
func (hh *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)

	data, dataName, err := rw.UploadBuffer()
	if err != nil {
		rw.JSONError(err)
		return
	}
	lg := logger.Gist(r.Context())
	lg.Info(dataName[0])

	rw.Bytes(data[dataName[0]].Bytes(), dataName[0], http.StatusOK)
}

// Post загрузка файла на сервер
// @Tags Data
// @Summary загрузка файла на сервер
// @Param file formData file true "загружаемый файл"
// @Accept mpfd
// @Success 200 {object} mdsungora.Minio "информация о загрузке"
// @Router /api/sun/data/upload [post]
func (hh *Handler) Post(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)

	obj := &mdsungora.User{}
	if err := rw.JSONBodyDecode(obj); err != nil {
		rw.JSONError(err)
		return
	}

	if err := obj.Insert(r.Context(), hh.st.DB(), boil.Infer()); err != nil {
		rw.JSONError(errs.NewBadRequest(err))
		return
	}

	rw.JSON(obj)
}
