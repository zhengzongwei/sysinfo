package commands

import (
	"Golang/tables"
	"fmt"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/spf13/cobra"
	"os"
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

		tables.CPUInfoTable()
	},
}

var memCmd = &cobra.Command{
	Use:   "mem",
	Short: "Query memory information",
	Run: func(cmd *cobra.Command, args []string) {
		memInfo, _ := mem.VirtualMemory()
		fmt.Printf("Memory Info: %+v\n", memInfo)
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
	// Add more subcommands here
}
