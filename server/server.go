package server

import (
	stdContext "context"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/YuShuanHsieh/trello-transform/errors"
)

type Server struct {
	ctx           stdContext.Context
	engine        *gin.Engine
	configuration *ServerConfiguration
}

func Default() *Server {
	config := defaultConfiguration()
	return New(config)
}

func New(config *ServerConfiguration) *Server {
	return &Server{
		ctx:           stdContext.Background(),
		engine:        gin.Default(),
		configuration: config,
	}
}

func (s *Server) Run() {
	s.engine.GET("/server/stop", s.stopServerHandler)
	s.engine.POST("/transform", s.transformHandler)
	err := s.engine.Run(fmt.Sprintf(":%d", s.configuration.Port))
	if err != nil {
		errors.Log(err.Error())
	}
}
