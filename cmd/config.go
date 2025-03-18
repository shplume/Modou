package cmd

import (
	"fmt"

	"github.com/shplume/Modou/core"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "config command line tool to run config",
	Long:  `A Fast and Flexible Config Toolkit.`,
	Run: func(cmd *cobra.Command, args []string) {
		server := core.ServerInstance
		server.Init()

		fmt.Println(core.ServerInstance.Config.GetAllConfigs())
	},
}

func init() {
	serverCmd.AddCommand(configCmd)
}
