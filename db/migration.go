package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/yashipro13/queryMaster/config"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

func createMigration(conf *config.Config) *migrate.Migrate {
	log.Printf("Connection String == %s", conf.GetDatabaseMigrationConnectionString())
	db, err := sql.Open("postgres", conf.GetDatabaseMigrationConnectionString())
	if err != nil {
		log.Fatalf("Failed to create Migration. Error = %s", err.Error())
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to create Migration with Instance. Error = %s", err.Error())
	}
	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "postgres", driver)
	if err != nil {
		log.Fatalf("Failed to create Migration with DB Instance. Error = %s", err.Error())
	}
	return m
}

func MigrateUp(conf *config.Config) error {
	m := createMigration(conf)
	defer m.Close()
	err := m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Printf("No migration change: %s", err.Error())
			return nil
		}
		return err
	}
	log.Print("Migration successful")
	return nil
}

func MigrateDown(conf *config.Config) error {
	m := createMigration(conf)
	err := m.Steps(-1)
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Printf("No migration change: %s", err.Error())
			return nil
		}
		return err
	}
	log.Print("Rollback successful")
	return nil
}
