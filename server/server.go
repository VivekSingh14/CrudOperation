package server

import (
	"fmt"
	"net/http"

	config "github.com/CrudOperation/configs"
	"github.com/gorilla/mux"
)

//adding comment
type Server struct {
	config *config.Config
	router *mux.Router
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *Server) ListenAndServe() {

	s.registerHandler()
	s.startServer()
}

func (s *Server) startServer() {

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", s.config.Server.Port),
		Handler: s.router,
	}

	srv.ListenAndServe()
}
