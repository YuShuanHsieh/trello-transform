package server

import (
	stdContext "context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/YuShuanHsieh/trello-transform/errors"
)

type Server struct {
	stop          chan os.Signal
	ctx           stdContext.Context
	engine        *gin.Engine
	configuration *ServerConfiguration
}

func Default() *Server {
	config := defaultConfiguration()
	return New(config)
}

func New(config *ServerConfiguration) *Server {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	return &Server{
		stop:          stop,
		ctx:           stdContext.Background(),
		engine:        gin.Default(),
		configuration: config,
	}
}

func (s *Server) Run() {
	s.engine.GET("/server/stop", s.stopServerHandler)
	s.engine.POST("/transform", s.transformHandler)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.configuration.Port),
		Handler: s.engine,
	}

	go func() {
		<-s.stop
		log.Println("Server is stopping")
		if err := srv.Shutdown(s.ctx); err != nil {
			log.Fatal(err)
		}
	}()

	if err := srv.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server stopped")
		} else {
			errors.Log(err.Error())
		}
	}
}
