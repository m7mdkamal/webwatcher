package cmd

import (
	"fmt"
	"os"

	"github.com/m7mdkamal/webwatcher/database"
	"github.com/spf13/cobra"
)

var DB database.Database

var rootCmd = &cobra.Command{
	Use:   "webwatcher",
	Short: "Webwatcher is app for watching new updates in the web",
	Long:  "Webwatcher is app for watching new updates in the web",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func init() {
	// cobra.OnInitialize(initConfig)
}
