package server

import (
	stdContext "context"
	"fmt"

	"github.com/kataras/iris"
)

type Server struct {
	ctx           stdContext.Context
	server        *iris.Application
	configuration *ServerConfiguration
}

func Default() *Server {
	config := defaultConfiguration()
	return New(config)
}

func New(config *ServerConfiguration) *Server {
	return &Server{
		ctx:           stdContext.Background(),
		server:        iris.Default(),
		configuration: config,
	}
}

func (s *Server) Run() {
	s.server.Get("/server/stop", s.stopServerHandler)
	s.server.Post("/transform", s.transformHandler)
	s.server.Run(iris.Addr(fmt.Sprintf(":%d", s.configuration.Port)))
}
