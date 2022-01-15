package general

import (
	"net/http"

	"sungora/app/config"
	"sungora/lib/logger"
	"sungora/lib/response"
	"sungora/lib/tpl"

	"github.com/shopspring/decimal"
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

// PageIndex пример динамического контента (html)
// @Tags General
// @Summary пример динамического контента (html)
// @Success 200 {string} string "html страница"
// @Router /index.hml [get]
func (hh *Handler) PageIndex(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)
	if r.URL.Path == "/" {
		r.URL.Path = response.IndexHtml
	}

	goods := Goods{
		{ID: 37, Name: "Item 10", Price: decimal.NewFromFloat(23.76)},
		{ID: 49, Name: "Item 2", Price: decimal.NewFromFloat(87.42)},
		{ID: 54, Name: "Item 30", Price: decimal.NewFromFloat(38.23)},
	}

	variable := map[string]interface{}{
		"Title": "PageIndex",
		"Goods": goods,
	}

	ret, err := tpl.ExecuteStorage(response.IndexHtml, variable)
	if err != nil {
		logger.Gist(r.Context()).Error(err)
		rw.Bytes([]byte("ошибка шаблона"), response.IndexHtml, http.StatusBadRequest)
		return
	}

	rw.Bytes(ret.Bytes(), r.URL.Path, http.StatusOK)
}

// PagePage пример динамического контента (html)
// @Tags General
// @Summary пример динамического контента (html)
// @Success 200 {string} string "html страница"
// @Router /page/index.html [get]
func (hh *Handler) PagePage(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)

	goods := Goods{
		{ID: 37, Name: "Item 10", Price: decimal.NewFromFloat(23.76)},
		{ID: 49, Name: "Item 2", Price: decimal.NewFromFloat(87.42)},
		{ID: 54, Name: "Item 30", Price: decimal.NewFromFloat(38.23)},
	}

	variable := map[string]interface{}{
		"Title": "PagePage",
		"Goods": goods,
	}

	ret, err := tpl.ExecuteStorage("page/page.html", variable)
	if err != nil {
		logger.Gist(r.Context()).Error(err)
		rw.Bytes([]byte("ошибка шаблона"), response.IndexHtml, http.StatusBadRequest)
		return
	}

	rw.Bytes(ret.Bytes(), r.URL.Path, http.StatusOK)
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

	rw.Bytes(data[dataName[0]].Bytes(), dataName[0], http.StatusOK)
}
