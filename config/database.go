package config

import "os"

func getDatabaseConfig() Database {
	return Database{
		Host:       os.Getenv("DATABASE_HOST"),
		Type:       "postgres",
		Username:   os.Getenv("DATABASE_USERNAME"),
		Password:   os.Getenv("DATABASE_PASSWORD"),
		Name:       os.Getenv("DATABASE_NAME"),
		SSLEnabled: os.Getenv("DATABASE_SSL_ENABLED") == "true",
	}
}

type Database struct {
	Host       string
	Type       string
	Username   string
	Password   string
	Name       string
	SSLEnabled bool
}
