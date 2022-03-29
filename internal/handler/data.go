package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"

	"sample/internal/model"
	"sample/lib/app/response"
	"sample/lib/logger"
	"sample/lib/minio"
	"sample/lib/storage/stpg"
	"sample/services/mdsample"
)

// UploadFile загрузка файла на сервер
// @Tags Data
// @Summary загрузка файла на сервер
// @Param file formData file true "загружаемый файл"
// @Accept mpfd
// @Produce octet-stream
// @Success 200 {file} file "file"
// @Router /sun/api/v1/data/upload-test [post]
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
// @Success 200 {object} mdsample.MinioSlice "информация о загрузке"
// @Router /sun/api/v1/data/upload [post]
func (hh *Handler) Upload(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)

	stM := model.NewData(stpg.Gist(), "")
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
// @Router /sun/api/v1/data/download/{id} [get]
func (hh *Handler) Download(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)
	id := uuid.MustParse(chi.URLParam(r, "id"))

	st, err := mdsample.FindMinio(r.Context(), hh.st.DB(), id)
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
