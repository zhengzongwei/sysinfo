package commands

import (
	"github.com/spf13/cobra"
	"sysinfo/tables"
)

var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "Query CPU information",
	Run: func(cmd *cobra.Command, args []string) {

		tables.CPUTable()
	},
}
