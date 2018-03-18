package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(watchersCmd)
	watchersCmd.AddCommand(listWatchersCmd)
	watchersCmd.AddCommand(clearWatchersCmd)
}

var watchersCmd = &cobra.Command{
	Use:   "watchers",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var listWatchersCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("reddit")
	},
}

var clearWatchersCmd = &cobra.Command{
	Use:   "clear",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CLEAR WATCHERS HERE")
	},
}
