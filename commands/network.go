package commands

import (
	"github.com/spf13/cobra"
	"sysinfo/tables"
)

var netCmd = &cobra.Command{
	Use:   "net",
	Short: "Query NET information",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var netShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show NET information",
	Run: func(cmd *cobra.Command, args []string) {
		tables.NetTable()
	},
}

func init() {
	netCmd.AddCommand(netShowCmd)
}
