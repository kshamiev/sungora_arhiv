package session

import (
	"time"
)

// TODO Mutex
type Session struct {
	t    time.Time
	data map[string]interface{}
}

func (s *Session) Get(key string) interface{} {
	if _, ok := s.data[key]; ok {
		return s.data[key]
	}
	return nil
}

func (s *Session) Set(key string, value interface{}) {
	s.data[key] = value
}

func (s *Session) Del(key string) {
	delete(s.data, key)
}
