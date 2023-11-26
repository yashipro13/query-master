package config

import "os"

type Elastic struct {
	Host string
}

func getElasticConfig() Elastic {
	return Elastic{
		Host: os.Getenv("ELASTIC_HOST"),
	}
}
