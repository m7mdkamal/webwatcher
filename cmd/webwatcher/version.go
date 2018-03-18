package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of webwatcher",
	Long:  "Print the version of webwatcher",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Webwatcher version is 0.0.0")
	},
}
