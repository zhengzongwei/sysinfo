package commands

import (
	"github.com/spf13/cobra"
	"sysinfo/tables"
)

var memCmd = &cobra.Command{
	Use:   "mem",
	Short: "Query memory information",
	Run: func(cmd *cobra.Command, args []string) {
		tables.MEMTable()
	},
}
