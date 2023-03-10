package server

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

type Server struct {
	http.Server
}

func NewServer(e *echo.Echo) *Server {
	return &Server{
		http.Server{
			Addr:              ":5000",
			Handler:           e,
			ReadTimeout:       30 * time.Second,
			ReadHeaderTimeout: 30 * time.Second,
			WriteTimeout:      30 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	log.Println("start serving in :5000")
	return s.ListenAndServe()
}
