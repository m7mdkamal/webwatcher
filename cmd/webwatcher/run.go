package cmd

import (
	"log"
	"sync"

	"github.com/jasonlvhit/gocron"
	"github.com/m7mdkamal/webwatcher/model"
	w "github.com/m7mdkamal/webwatcher/watcher"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run tasks",
	Long:  "Run tasks",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func runTasks() {
	var wg sync.WaitGroup

	tasks, err := DB.GetTasks()
	if err != nil {
		panic(err)
	}
	for _, task := range tasks {
		s := gocron.NewScheduler()
		s.Every(uint64(task.Interval)).Seconds().Do(single, task)
		go s.Start()
		wg.Add(1)
	}
}

func single(task model.Task) {
	log.Printf("[%s] watcher started", task.Name)
	watcher, err := DB.GetWatcherByTask(&task)
	if err != nil {
		panic(err)
	}
	results := w.NewWatcher(watcher.ID, task.Filter, task.Parameters).Run()
	log.Printf("[%s] watcher found %d", task.Name, len(results))
	for _, result := range results {
		log.Printf("[%s] watcher found: %s", task.Name, result.Title)
		result.TaskID = task.ID
		DB.CreateResult(&result)
	}
	// err := beeep.Notify("Title", "Message body", "assets/information.png")
	// if err != nil {
	// 	panic(err)
	// }
	log.Printf("[%s] watcher ended. Should start after %d secs", task.Name, task.Interval)
	// wg.Done()
}
