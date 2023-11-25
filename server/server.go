package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yashipro13/queryMaster/config"
	"github.com/yashipro13/queryMaster/repository"
	"github.com/yashipro13/queryMaster/users"
	"log"
)

type server struct {
	router *gin.Engine
	config *config.Config
}

func New() (*server, error) {
	cfg := config.NewConfig()
	conn, err := createDBConn(cfg.GetDatabaseConnectionString())
	if err != nil {
		return nil, err
	}
	db, err := repository.NewRepo(conn)
	if err != nil {
		log.Printf("failed to initialize repository")
		return nil, err
	}
	router := NewRouter(users.Service{DBManager: db})
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

func createDBConn(connString string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	return pool, err
}
