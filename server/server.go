package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yashipro13/queryMaster/config"
	"github.com/yashipro13/queryMaster/hashtags"
	"github.com/yashipro13/queryMaster/ingestor"
	"github.com/yashipro13/queryMaster/repository"
	"github.com/yashipro13/queryMaster/users"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"time"
)

type server struct {
	router   *gin.Engine
	config   *config.Config
	repo     *repository.Repo
	esClient *elastic.Client
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
	es, err := initElasticsearch(cfg.ElasticConfigs().Host)
	if err != nil {
		log.Printf("failed to initialize es client with err %s", err.Error())
		return nil, err
	}
	router := NewRouter(users.Service{DBManager: db}, hashtags.Service{DBManager: db})
	server := &server{
		router:   router,
		config:   cfg,
		repo:     db,
		esClient: es,
	}
	return server, nil
}

func (s server) Run() {
	fmt.Printf("Running on port %d", s.config.AppPort())
	go func() { s.syncData(context.Background()) }()
	s.router.Run(fmt.Sprintf(":%d", s.config.AppPort()))
}

func createDBConn(connString string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	return pool, err
}

func initElasticsearch(host string) (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL(host))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (s server) syncData(ctx context.Context) {
	ingestionSvc := ingestor.Service{
		DBManager:     s.repo,
		ElasticClient: s.esClient,
	}
	log.Printf("attempting to ingest data")
	for {
		svcErr := ingestionSvc.IngestData(ctx)
		if svcErr != nil {
			log.Printf("failed ingestion of data with error %s", svcErr.Message)
		}
		time.Sleep(time.Second * 10)
	}
}
