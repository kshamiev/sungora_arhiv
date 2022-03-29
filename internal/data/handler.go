package data

import (
	"net/http"

	"sungora/lib/app/response"
	"sungora/lib/logger"
	"sungora/lib/minio"
	"sungora/lib/storage"
	"sungora/lib/storage/stpg"
	"sungora/services/mdsungora"

	"github.com/google/uuid"

	"github.com/go-chi/chi"
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
// @Success 200 {file} file "file"
// @Router /api/sun/data/upload-test [post]
func (hh *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)

	data, dataName, err := rw.UploadBuffer()
	if err != nil {
		rw.JSON(err)
		return
	}
	lg := logger.Get(r.Context())
	lg.Info(dataName[0])

	rw.Bytes(data[dataName[0]].Bytes(), dataName[0], http.StatusOK)
}

// Upload загрузка файла на сервер (minio)
// @Tags Data
// @Summary загрузка файла на сервер (minio)
// @Param file formData file true "загружаемый файл"
// @Accept mpfd
// @Success 200 {object} mdsungora.MinioSlice "информация о загрузке"
// @Router /api/sun/data/upload [post]
func (hh *Handler) Upload(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)

	stM := NewModel(stpg.Gist(), "")
	res, err := stM.UploadRequest(rw, "test")
	if err != nil {
		rw.JSON(err)
		return
	}

	rw.JSON(res)
}

// Download загрузка файла c сервера (minio)
// @Tags Data
// @Summary загрузка файла с сервера (minio)
// @Param id path string true "ID"
// @Produce octet-stream
// @Success 200 {file} file "file"
// @Router /api/sun/data/download/{id} [get]
func (hh *Handler) Download(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)
	id := uuid.MustParse(chi.URLParam(r, "id"))

	st, err := mdsungora.FindMinio(r.Context(), hh.st.DB(), id)
	if err != nil {
		rw.JSON(err)
		return
	}

	res, err := minio.GetFile(r.Context(), st.Bucket, st.ID.String())
	if err != nil {
		rw.JSON(err)
		return
	}

	rw.Reader(res, st.Name, http.StatusOK)
}
