package commands

import (
	"fmt"
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

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCMD.AddCommand(cpuCmd)
	rootCMD.AddCommand(diskCmd)
	rootCMD.AddCommand(memCmd)
	rootCMD.AddCommand(netCmd)
	rootCMD.AddCommand(hostCmd)
}
