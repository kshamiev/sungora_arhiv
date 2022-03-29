package session

import (
	"net/http"
	"sync"
	"time"

	"sungora/lib/app/response"

	"github.com/google/uuid"
)

type SessionBus map[string]*Session

func NewSessionBus(sessionTimeout time.Duration) SessionBus {
	var sessions = make(SessionBus)
	go sessions.controlSessionBus(sessionTimeout)
	return sessions
}

// controlSessionBus жизненный цикл сессии
func (bus SessionBus) controlSessionBus(sessionTimeout time.Duration) {
	for {
		time.Sleep(time.Minute)
		for key := range bus {
			if sessionTimeout < time.Since(bus[key].t) {
				delete(bus, key)
			}
		}
	}
}

var mu sync.Mutex

func (bus SessionBus) Get(key string) interface{} {
	mu.Lock()
	defer mu.Unlock()
	if elm, ok := bus[key]; ok {
		if _, ok := elm.data[key]; ok {
			return elm.data[key]
		}
	}
	return nil
}

func (bus SessionBus) Set(key string, val interface{}) {
	elm := new(Session)
	elm.t = time.Now()
	elm.data = map[string]interface{}{
		key: val,
	}
	mu.Lock()
	bus[key] = elm
	mu.Unlock()
}

func (bus SessionBus) Del(key string) {
	mu.Lock()
	delete(bus, key)
	mu.Unlock()
}

// GetSessionCookie Получение сессии по куке пришедшей из запроса
func (bus SessionBus) GetSessionCookie(
	r *http.Request, w http.ResponseWriter, cookieName string, path []string) *Session {
	rw := response.New(r, w)
	token := rw.CookieGet(cookieName)
	if token == "" {
		token = uuid.New().String()
		rw.CookieSet(cookieName, token, path)
	}
	return bus.GetSession(token)
}

// GetSession Получение сессии по токену
func (bus SessionBus) GetSession(token string) *Session {
	mu.Lock()
	defer mu.Unlock()
	if elm, ok := bus[token]; ok {
		elm.t = time.Now()
		return elm
	}
	bus[token] = &Session{
		t:    time.Time{},
		data: make(map[string]interface{}),
	}
	return bus[token]
}
