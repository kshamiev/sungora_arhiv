package web

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

// конфигурация HTTP
type HttpServerConfig struct {
	Proto          string        `yaml:"proto"`          // Server Proto
	Host           string        `yaml:"host"`           // Server Host
	Port           int           `yaml:"port"`           // Server Port
	ReadTimeout    time.Duration `yaml:"readTimeout"`    // Время ожидания web запроса в секундах
	WriteTimeout   time.Duration `yaml:"writeTimeout"`   // Время ожидания окончания передачи ответа в секундах
	RequestTimeout time.Duration `yaml:"requestTimeout"` // Время ожидания окончания выполнения запроса
	IdleTimeout    time.Duration `yaml:"idleTimeout"`    // Время ожидания следующего запроса
	MaxHeaderBytes int           `yaml:"maxHeaderBytes"` // Максимальный размер заголовка получаемого от браузера клиента в байтах
}

// сервер http(s)
type HttpServer struct {
	server    *http.Server  // сервер HTTP
	chControl chan struct{} // управление ожиданием завершения работы сервера
	lis       net.Listener
}

// создание и старт вебсервера (HTTP(S))
func NewServer(cfg *HttpServerConfig, mux http.Handler) (comp *HttpServer, err error) {
	comp = &HttpServer{
		server: &http.Server{
			Addr:           fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Handler:        mux,
			ReadTimeout:    cfg.ReadTimeout,
			WriteTimeout:   cfg.WriteTimeout,
			IdleTimeout:    cfg.IdleTimeout,
			MaxHeaderBytes: cfg.MaxHeaderBytes,
		},
		chControl: make(chan struct{}),
	}

	if comp.lis, err = net.Listen("tcp", comp.server.Addr); err != nil {
		return
	}

	go func() {
		_ = comp.server.Serve(comp.lis)
		close(comp.chControl)
	}()

	return comp, nil
}

// завершение работы сервера (HTTP(S))
func (comp *HttpServer) CloseWait() {
	if comp.lis == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := comp.server.Shutdown(ctx); err != nil {
		return
	}

	if err := comp.lis.Close(); err != nil {
		return
	}

	<-comp.chControl
}

// получение обработчика запросов
func (comp *HttpServer) GetRoute() *chi.Mux {
	if _, ok := comp.server.Handler.(*chi.Mux); ok {
		return comp.server.Handler.(*chi.Mux)
	}
	return nil
}
