package main

import (
	"github.com/spf13/cobra"
	"github.com/yashipro13/queryMaster/command"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "query-master",
	Short: "QueryMaster",
}

func main() {
	setDefaultCommandIfNonePresent()
	rootCmd.AddCommand(command.Server())
	rootCmd.AddCommand(command.Migration())
	rootCmd.Execute()
}

func setDefaultCommandIfNonePresent() {
	if len(os.Args) < 2 {
		os.Args = append([]string{os.Args[0], "server"})
	}
}
