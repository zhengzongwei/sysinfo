package commands

import (
	"github.com/spf13/cobra"
	"sysinfo/tables"
)

var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Query HOST information",
	Run: func(cmd *cobra.Command, args []string) {

		tables.HostTable()
	},
}
