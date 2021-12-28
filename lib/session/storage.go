package session

import (
	"net/http"
	"time"

	"sungora/lib/response"
	"sungora/lib/typ"
)

// TODO Mutex
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

func (bus SessionBus) Get(key string) interface{} {
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
	bus[key] = elm
}

func (bus SessionBus) Del(key string) {
	delete(bus, key)
}

// GetSessionCookie Получение сессии по куке пришедшей из запроса
func (bus SessionBus) GetSessionCookie(
	r *http.Request, w http.ResponseWriter, cookieName string, path []string) *Session {
	rw := response.New(r, w)
	token := rw.CookieGet(cookieName)
	if token == "" {
		token = typ.UUIDNew().String()
		rw.CookieSet(cookieName, token, path)
	}
	return bus.GetSession(token)
}

// GetSession Получение сессии по токену
func (bus SessionBus) GetSession(token string) *Session {
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
