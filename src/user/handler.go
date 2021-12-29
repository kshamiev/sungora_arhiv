package user

import (
	"net/http"

	"sungora/lib/errs"
	"sungora/lib/logger"
	"sungora/lib/response"
	"sungora/lib/storage/pgsql"
	"sungora/lib/typ"
	"sungora/services/mdsun"
	"sungora/src/client"

	"github.com/go-chi/chi"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Handler struct {
	db *pgsql.Storage
}

func NewHandler() *Handler {
	return &Handler{
		db: pgsql.Gist(),
	}
}

// GetSlice
// @Tags User
// @Summary Получение списка пользователей
// @Success 200 {array} mdsun.UserSlice ""
// @Router /api/sun/users [get]
func (hh *Handler) GetSlice(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)

	list, err := mdsun.Users(
		qm.OrderBy(mdsun.UserColumns.CreatedAt+" ASC"),
		qm.Offset(0), qm.Limit(20),
	).All(r.Context(), hh.db.DB())
	if err != nil {
		rw.JSONError(errs.NewBadRequest(err))
		return
	}

	rw.JSON(list)
}

// Get
// @Tags User
// @Summary Получение пользователя
// @Param id path string true "ID"
// @Success 200 {object} mdsun.User ""
// @Router /api/sun/user/{id} [get]
func (hh *Handler) Get(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)

	obj, err := mdsun.FindUser(r.Context(), hh.db.DB(), typ.UUIDMustParse(chi.URLParam(r, "id")))
	if err != nil {
		rw.JSONError(errs.NewBadRequest(err))
		return
	}

	rw.JSON(obj)
}

// Post
// @Tags User
// @Summary Создание пользователя
// @Param data body mdsun.User true "пользователь"
// @Success 200 {object} mdsun.User ""
// @Router /api/sun/user/{id} [post]
func (hh *Handler) Post(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)

	obj := &mdsun.User{}
	if err := rw.JSONBodyDecode(obj); err != nil {
		rw.JSONError(err)
		return
	}

	if err := obj.Insert(r.Context(), hh.db.DB(), boil.Infer()); err != nil {
		rw.JSONError(errs.NewBadRequest(err))
		return
	}

	rw.JSON(obj)
}

// Put
// @Tags User
// @Summary Изменение пользователя
// @Param id path string true "ID"
// @Param data body mdsun.User true "пользователь"
// @Success 200 {object} mdsun.User ""
// @Router /api/sun/user/{id} [put]
func (hh *Handler) Put(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)

	obj := &mdsun.User{}
	if err := rw.JSONBodyDecode(obj); err != nil {
		rw.JSONError(err)
		return
	}

	if _, err := obj.Update(r.Context(), hh.db.DB(), boil.Infer()); err != nil {
		rw.JSONError(errs.NewBadRequest(err))
		return
	}

	rw.JSON(obj)
}

// Delete
// @Tags User
// @Summary Удаление пользователя
// @Param id path string true "ID"
// @Success 200 {string} string "OK"
// @Router /api/sun/user/{id} [delete]
func (hh *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)

	obj := &mdsun.User{ID: typ.UUIDMustParse(chi.URLParam(r, "id"))}

	if _, err := obj.Delete(r.Context(), hh.db.DB()); err != nil {
		rw.JSONError(errs.NewBadRequest(err))
		return
	}

	rw.JSON("OK")
}

// Test
// @Tags User
// @Summary Тестовый обработчик для примера
// @Description Тестовый обработчик для примера
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} mdsun.User "пользователь"
// @Failure 400 {string} response.Data "отрицательный ответ"
// @Security ApiKeyAuth
// @Router /api/sun/user-test/{id} [get]
// @Deprecated
// Deprecated
func (hh *Handler) Test(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)

	usM := NewModel(hh.db)
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
