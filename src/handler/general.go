package handler

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gorilla/websocket"

	"sungora/lib/errs"
	"sungora/lib/logger"
	"sungora/lib/response"
	"sungora/lib/uuid"
	"sungora/lib/web"
	"sungora/src/config"
	"sungora/src/model"
	"sungora/src/service"
)

type General struct {
	WsBus web.SocketBus
}

// общие запросы
func NewGeneral() *General {
	return &General{
		WsBus: web.NewSocketBus(),
	}
}

// Ping ping
// @Summary ping
// @Tags General
// @Router /api/v1/general/ping [get]
// @Success 200 {string} string "OK"
func (c *General) Ping(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)
	rw.JSON("OK")
}

// GetVersion получение версии приложения
// @Summary получение версии приложения
// @Tags General
// @Router /api/v1/general/version [get]
// @Success 200 {string} string "version"
func (c *General) GetVersion(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)
	rw.JSON(config.Get().App.Version)
}

// Test test
// @Summary test
// @Tags General
// @Router /api/v1/general/test/{id} [get]
// @Param id path string true "ID"
// @Success 200 {object} typ.Users "user"
func (c *General) Test(w http.ResponseWriter, r *http.Request) {
	rw := response.New(r, w)

	usM := model.NewUser()
	us, err := usM.Load(r.Context(), uuid.UUIDMustParse(chi.URLParam(r, "id")))
	if err != nil {
		rw.JSONError(err)
		return
	}

	cli := service.GetSampleClient()
	if _, err := cli.GetVersion(r.Context(), &empty.Empty{}); err != nil {
		rw.JSONError(err)
		return
	}

	rw.JSON(us)
}

// @Summary пример работы с вебсокетом (http://localhost:8080/template/gorilla/index.html)
// @Tags General
// @Router /api/v1/websocket/gorilla/{id} [get]
// @Success 101 {string} string "Switching Protocols to websocket"
// @Security ApiKeyAuth
func (c *General) GetWebSocketSample(w http.ResponseWriter, r *http.Request) {
	var (
		ws         *websocket.Conn
		wsResponse http.Header

		err error
	)
	rw := response.New(r, w)
	lg := logger.GetLogger(r.Context())

	// переключаемся на вебсокет
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	if ws, err = upgrader.Upgrade(w, r, wsResponse); err != nil {
		rw.JSONError(errs.NewBadRequest(err, "не удалось переключить протокол на ws"))
		return
	}
	defer func() {
		if err := ws.Close(); err != nil {
			lg.WithError(err).Error("WS close connect error")
		} else {
			lg.Info("WS close connect ok")
		}
	}()

	// создаем  и запускаем клиента
	client := &chatWS{
		Ws:  ws,
		Ctx: r.Context(),
	}
	c.WsBus.StartClient(chi.URLParam(r, "id"), client)
}

// пример обработчика клиента
type chatWS struct {
	Ws  *websocket.Conn
	Ctx context.Context
}

// HookStartClient метод при подключении и старте нового пользователя
func (cli *chatWS) HookStartClient(cntClient int) error {
	logger.GetLogger(cli.Ctx).Info("WS hook StartClient: ", cntClient)
	return nil
}

// HookGetMessage метод при получении данных из вебсокета пользователя
func (cli *chatWS) HookGetMessage(cntClient int) (interface{}, error) {
	lg := logger.GetLogger(cli.Ctx)
	msg := &model.Message{}
	if err := cli.Ws.ReadJSON(msg); err != nil {
		return nil, err
	}
	msg.Author += " - WS OK"
	lg.Info("WS hook GetMessage")
	if err := ioutil.WriteFile(msg.FileName, msg.FileData, 0666); err != nil {
		return nil, err
	}
	return msg, nil
}

// HookSendMessage метод при отправке данных пользователю
func (cli *chatWS) HookSendMessage(msg interface{}, cntClient int) error {
	lg := logger.GetLogger(cli.Ctx)
	res := &model.Message{}
	switch o := msg.(type) {
	case *model.Message:
		res = o
	case response.Error:
		lg.WithError(o).Error(o.Response())
		res.Author = "system 1"
		_ = res.Body.Marshal(o.Response())
	case error:
		lg.WithError(o).Error("Other (unexpected) error")
		res.Author = "system 2"
		_ = res.Body.Marshal("Other (unexpected) error")
	}
	_ = cli.Ws.WriteJSON(res)
	lg.Info("WS hook SendMessage")
	return nil
}

// Ping проверка соединения с пользователем
func (cli *chatWS) Ping() error {
	logger.GetLogger(cli.Ctx).Info("WS hook Ping")
	return cli.Ws.WriteMessage(websocket.PingMessage, []byte{})
}
