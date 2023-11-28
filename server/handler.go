package server

import "github.com/CrudOperation/handler"

func (s *Server) registerHandler() {
	handler.RegisterHandlers(s.router, s.config)
}
