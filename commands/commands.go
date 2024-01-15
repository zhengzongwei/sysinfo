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

var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Query HOST information",
	Run: func(cmd *cobra.Command, args []string) {

		tables.HostTable()
	},
}

var netCmd = &cobra.Command{
	Use:   "net",
	Short: "Query NET information",
	Run: func(cmd *cobra.Command, args []string) {

		tables.NetTable()
	},
}

var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Query NET information",
	Run: func(cmd *cobra.Command, args []string) {

		tables.DiskTable()
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
	rootCMD.AddCommand(cpuCmd)
	rootCMD.AddCommand(memCmd)
	rootCMD.AddCommand(hostCmd)
	rootCMD.AddCommand(netCmd)
	rootCMD.AddCommand(diskCmd)

	// Add more subcommands here

	//
	//rootCMD.SortCommands = true
}
