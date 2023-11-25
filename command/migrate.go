package command

import (
	"github.com/spf13/cobra"
	"github.com/yashipro13/queryMaster/config"
	"github.com/yashipro13/queryMaster/db"
)

func Migration() *cobra.Command {
	mg := &cobra.Command{
		Use:   "migrate",
		Short: "migration for database",
	}
	mg.AddCommand(migrateUpCommand())
	mg.AddCommand(migrateDownCommand())
	return mg
}

func migrateUpCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "up migration for database",
		RunE: func(cmd *cobra.Command, args []string) error {
			conf := config.NewConfig()
			return db.MigrateUp(conf)
		},
	}
}

func migrateDownCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "down",
		Short: "down migration for database",
		RunE: func(cmd *cobra.Command, args []string) error {
			conf := config.NewConfig()
			return db.MigrateDown(conf)
		},
	}
}
