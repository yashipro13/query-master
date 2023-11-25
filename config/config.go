package config

import (
	"os"
	"strconv"
)

type Config struct {
	database Database
	elastic  Elastic
	appPort  int
}

func NewConfig() *Config {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}
	return &Config{
		database: getDatabaseConfig(),
		elastic:  getElasticConfig(),
		appPort:  port,
	}
}

func (c *Config) DatabaseConfigs() Database {
	return c.database
}

func (c *Config) ElasticConfigs() Elastic {
	return c.elastic
}

func (c *Config) AppPort() int {
	return c.appPort
}
