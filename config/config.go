package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	database Database
	elastic  Elastic
	appPort  int
}

func (c *Config) GetDatabaseConnectionString() string {
	return fmt.Sprintf("%s://%s:%s@%s/%s", c.database.Type, c.database.Username, c.database.Password, c.database.Host, c.database.Name)
}

func (c *Config) GetDatabaseMigrationConnectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password='%s' host=%s port=%d sslmode=disable", c.database.Name, c.database.Username, c.database.Password, c.database.Host, 5432)
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
