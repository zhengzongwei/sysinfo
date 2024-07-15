package commands

import (
	"github.com/spf13/cobra"
	"sysinfo/tables"
)

var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Query NET information",
	Run: func(cmd *cobra.Command, args []string) {

		tables.DiskTable()
	},
}
