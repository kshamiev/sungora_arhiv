package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"sample/internal/client"
	"sample/internal/model"
	"sample/lib/app/response"
	"sample/lib/errs"
	"sample/lib/logger"
	"sample/services/mdsample"
	"sample/services/pbsample"
)

// GetSlice
// @Tags User
// @Summary Получение списка пользователей
// @Success 200 {array} mdsample.UserSlice ""
// @Router /sun/api/v1/users [get]
func (hh *Handler) GetSlice(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)

	list, err := mdsample.Users(
		qm.OrderBy(mdsample.UserColumns.CreatedAt+" ASC"),
		qm.Offset(0), qm.Limit(20),
	).All(r.Context(), hh.st.DB())
	if err != nil {
		rw.JSON(errs.New(err))
		return
	}

	rw.JSON(list)
}

// Get
// @Tags User
// @Summary Получение пользователя
// @Param id path int true "ID"
// @Success 200 {object} mdsample.User ""
// @Router /sun/api/v1/user/{id} [get]
func (hh *Handler) Get(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		rw.JSON(errs.New(err))
		return
	}

	obj, err := mdsample.FindUser(r.Context(), hh.st.DB(), id)
	if err != nil {
		rw.JSON(errs.New(err))
		return
	}

	rw.JSON(obj)
}

// Post
// @Tags User
// @Summary Создание пользователя
// @Param data body mdsample.User true "пользователь"
// @Success 200 {object} mdsample.User ""
// @Router /sun/api/v1/user/{id} [post]
func (hh *Handler) Post(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)

	obj := &mdsample.User{}
	if err := rw.JSONBodyDecode(obj); err != nil {
		rw.JSON(err)
		return
	}

	if err := obj.Insert(r.Context(), hh.st.DB(), boil.Infer()); err != nil {
		rw.JSON(errs.New(err))
		return
	}

	rw.JSON(obj)
}

// Put
// @Tags User
// @Summary Изменение пользователя
// @Param id path int true "ID"
// @Param data body mdsample.User true "пользователь"
// @Success 200 {object} mdsample.User ""
// @Router /sun/api/v1/user/{id} [put]
func (hh *Handler) Put(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)

	obj := &mdsample.User{}
	if err := rw.JSONBodyDecode(obj); err != nil {
		rw.JSON(err)
		return
	}

	if _, err := obj.Update(r.Context(), hh.st.DB(), boil.Infer()); err != nil {
		rw.JSON(errs.New(err))
		return
	}

	rw.JSON(obj)
}

// Delete
// @Tags User
// @Summary Удаление пользователя
// @Param id path int true "ID"
// @Success 200 {string} string "OK"
// @Router /sun/api/v1/user/{id} [delete]
func (hh *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		rw.JSON(errs.New(err))
		return
	}

	obj := &mdsample.User{ID: id}
	if _, err := obj.Delete(r.Context(), hh.st.DB()); err != nil {
		rw.JSON(errs.New(err))
		return
	}

	rw.JSON("OK")
}

// Sample
// @Tags User
// @Summary Тестовый обработчик для примера
// @Description Тестовый обработчик для примера
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} mdsample.User "пользователь"
// @Failure 400 {string} response.Data "отрицательный ответ"
// @Security ApiKeyAuth
// @Router /sun/api/v1/user-sample/{id} [get]
// @Deprecated
// Deprecated
func (hh *Handler) Sample(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		rw.JSON(errs.New(err))
		return
	}

	usM := model.NewUser()
	us, err := usM.Load(r.Context(), id)
	if err != nil {
		rw.JSON(err)
		return
	}

	lg := logger.Get(r.Context())
	lg.Info("User.Test")
	err = errors.New("sample error")
	err = errs.New(err, "user message error")
	lg.WithError(err).Error(err.(*errs.Errs).Response())

	cli := client.GistSampleGRPC()
	if _, err := cli.Ping(r.Context(), &pbsample.Test{Text: "Fantik"}); err != nil {
		rw.JSON(err)
		return
	}

	rw.JSON(us)
}
