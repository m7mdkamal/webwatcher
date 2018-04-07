package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(resetCmd)
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "remove webwatcher tasks and results",
	Long:  "remove webwatcher tasks and results",
	Run: func(cmd *cobra.Command, args []string) {
		resetResults()
		resetTasks()
	},
}

func resetTasks() {
	err := DB.DeleteTasks()
	if err != nil {
		panic(err)
	}
}

func resetResults() {
	err := DB.DeleteResults()
	if err != nil {
		panic(err)
	}
}
