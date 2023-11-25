package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yashipro13/queryMaster/config"
)

type server struct {
	router *gin.Engine
	config *config.Config
}

func New() (*server, error) {
	cfg := config.NewConfig()
	router := NewRouter()
	server := &server{
		router: router,
		config: cfg,
	}
	return server, nil
}

func (s server) Run() {
	fmt.Printf("Running on port %d", s.config.AppPort())
	s.router.Run(fmt.Sprintf(":%d", s.config.AppPort()))
}
