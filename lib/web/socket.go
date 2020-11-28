package web

import (
	"errors"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// интерфейс клиента для взаимодействия с вебсокетом
type SocketClient interface {
	// Вызывается единожды в момент подключения нового клиента (заходит на страницу с вебсокетом)
	HookStartClient(cntClient int) error
	// Читает сообщения вебсокета
	HookGetMessage(cntClient int) (interface{}, error)
	// Отправляет сообщения в вебсокет (централизованно вызывается для всех подключенных)
	HookSendMessage(msg interface{}, cntClient int) error
	// Проверка и пролонгация подключения
	Ping() error
}

var mu sync.Mutex

// шина обработчиков вебсокетов по идентификаторам
type SocketBus map[string]*SocketHandler

// создание шины для обработчиков
func NewSocketBus() SocketBus { return make(SocketBus) }

// обработчик клиентов
type SocketHandler struct {
	broadcast chan interface{}      // канал передачи данных всем клиентам обработчика
	clients   map[SocketClient]bool // массив всех клиентов обработчика
	isClose   bool                  // признак завершения работы обработчика
}

// GetWSHandler получение обработчика
func (bus SocketBus) GetWSHandler(wsbusID string) (*SocketHandler, error) {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := bus[wsbusID]; ok {
		return bus[wsbusID], nil
	}
	return nil, errors.New("not found ws handler")
}

// StartClient инициализация клиента по условному идентификатору
// инициализация и регистрация обработчика в шине
// регистрация и старт работы нового клиента в обратчике
// управление работой клиента
func (bus SocketBus) StartClient(wsbusID string, client SocketClient) {
	// инициализация обработчика в шине
	mu.Lock()
	room, ok := bus[wsbusID]
	if !ok {
		room = &SocketHandler{
			broadcast: make(chan interface{}),
			clients:   make(map[SocketClient]bool),
		}
		bus[wsbusID] = room
		go room.control()
	}
	mu.Unlock()

	// регистрация клиента и завершение работы
	room.clients[client] = true
	defer func() {
		delete(room.clients, client)
		if len(room.clients) == 0 {
			room.isClose = true
			delete(bus, wsbusID)
		}
	}()

	// старт работы клиента
	if err := client.HookStartClient(len(room.clients)); err != nil {
		_ = client.HookSendMessage(err, len(room.clients))
		return
	}

	// здесь мы лочимся и обрабатываем входящие сообщения до выхода
	for {
		msg, err := client.HookGetMessage(len(room.clients))
		if err != nil {
			if _, ok := err.(*websocket.CloseError); !ok {
				_ = client.HookSendMessage(err, len(room.clients))
			}
			return
		}
		if msg != nil {
			room.broadcast <- msg // посылаем всем подключенным пользователям
		}
	}
}

// control здесь мы отправляем сообщения всем подключенным клиентам и пингуем
func (room *SocketHandler) control() {
	ticker := time.NewTicker(time.Minute)
	for {
		if room.isClose {
			return
		}
		select {
		// проверка соединений с клиентами
		case <-ticker.C:
			for handler := range room.clients {
				// если достучаться до клиента не удалось, то удаляем его
				if err := handler.Ping(); err != nil {
					delete(room.clients, handler)
					continue
				}
			}
		// каждому зарегистрированному клиенту шлем сообщение
		case message := <-room.broadcast:
			for handler := range room.clients {
				_ = handler.HookSendMessage(message, len(room.clients))
			}
		}
	}
}

// SendMessage отправка сообщений всем покдлюченным клиентам
func (room *SocketHandler) SendMessage(msg interface{}) {
	if msg != nil {
		room.broadcast <- msg
	}
}

// GetClientCnt получение количества подключенных клиентов
func (room *SocketHandler) GetClientCnt() int {
	return len(room.clients)
}
