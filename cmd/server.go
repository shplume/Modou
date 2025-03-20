package cmd

import (
	"fmt"
	"os"

	"github.com/shplume/Modou/core"
	"github.com/shplume/Modou/router"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server command line tool to run server",
	Long:  `A Fast and Flexible Server Toolkit.`,
	Run: func(cmd *cobra.Command, args []string) {
		server := core.ServerInstance

		router.Init(server)

		server.Run()
	},
}

func Execute() {
	if err := serverCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
