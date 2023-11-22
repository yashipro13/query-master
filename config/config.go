package config

type config struct {
	database Database
	elastic  Elastic
}

var cfg config

func Load() {
	cfg = config{
		database: getDatabaseConfig(),
		elastic:  getElasticConfig(),
	}
}

func DatabaseConfigs() Database {
	return cfg.database
}

func ElasticConfigs() Elastic {
	return cfg.elastic
}
