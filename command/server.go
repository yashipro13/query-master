package command

import (
	"github.com/spf13/cobra"
	"github.com/yashipro13/queryMaster/server"

	"log"
)

func Server() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "starts the http server",
		Run: func(cmd *cobra.Command, args []string) {
			svr, err := server.New()
			if err != nil {
				log.Fatalf("Failed to create server. Err = %s", err.Error())
			}
			svr.Run()
		},
	}
	return cmd
}
