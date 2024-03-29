package handler

import (
	"context"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"

	"sample/internal/model"
	"sample/lib/app/response"
	"sample/lib/errs"
	"sample/lib/logger"
)

// WebSocketSample пример работы с веб-сокетом (http://localhost:8080/assets/gorilla/index.html)
// @Tags Websocket
// @Summary пример работы с веб-сокетом (http://localhost:8080/assets/gorilla/index.html)
// @Success 101 {string} string "Switching Protocols to websocket"
// @Security ApiKeyAuth
// @Router /sun/api/v1/websocket/gorilla/{id} [get]
func (hh *Handler) WebSocketSample(w http.ResponseWriter, r *http.Request) {
	var (
		ws         *websocket.Conn
		wsResponse http.Header

		err error
	)
	rw := response.New(r, w)
	lg := logger.Get(r.Context())

	// переключаемся на вебсокет
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	if ws, err = upgrader.Upgrade(w, r, wsResponse); err != nil {
		rw.JSON(errs.New(err, "не удалось переключить протокол на ws"))
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
	cli := &chatWS{
		Ws:  ws,
		Ctx: r.Context(),
	}
	hh.wsBus.StartClient(chi.URLParam(r, "id"), cli)
}

// chatWS пример обработчика клиента
type chatWS struct {
	Ws  *websocket.Conn
	Ctx context.Context
}

// HookStartClient метод при подключении и старте нового пользователя
func (cli *chatWS) HookStartClient(cntClient int) error {
	logger.Get(cli.Ctx).Info("WS hook StartClient: ", cntClient)
	return nil
}

// HookGetMessage метод при получении данных из вебсокета пользователя
func (cli *chatWS) HookGetMessage(cntClient int) (interface{}, error) {
	lg := logger.Get(cli.Ctx)
	msg := &model.Message{}
	if err := cli.Ws.ReadJSON(msg); err != nil {
		return nil, err
	}
	msg.Author += " - WS OK"
	lg.Info("WS hook GetMessage")
	if msg.FileSize > 0 {
		if err := os.WriteFile(msg.FileName, msg.FileData, 0o600); err != nil {
			return nil, err
		}
	}
	return msg, nil
}

// HookSendMessage метод при отправке данных пользователю
func (cli *chatWS) HookSendMessage(msg interface{}, cntClient int) error {
	lg := logger.Get(cli.Ctx)
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
	logger.Get(cli.Ctx).Info("WS hook Ping")
	return cli.Ws.WriteMessage(websocket.PingMessage, []byte{})
}
