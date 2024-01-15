package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"sysinfo/tables"
)

var rootCMD = &cobra.Command{
	Use:   "sysinfo",
	Short: "Query system information",
	Long:  "A command-line tool to query various system information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify a subcommand. See 'sysinfo --help' for more details.")
	},
}

var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Query HOST information",
	Run: func(cmd *cobra.Command, args []string) {

		tables.HostTable()
	},
}

var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "Query CPU information",
	Run: func(cmd *cobra.Command, args []string) {

		tables.CPUTable()
	},
}

var memCmd = &cobra.Command{
	Use:   "mem",
	Short: "Query memory information",
	Run: func(cmd *cobra.Command, args []string) {
		tables.MEMTable()
	},
}

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Register subcommands
	rootCMD.AddCommand(hostCmd)
	rootCMD.AddCommand(cpuCmd)
	rootCMD.AddCommand(memCmd)
	// Add more subcommands here
}
