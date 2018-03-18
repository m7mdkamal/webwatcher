package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/gen2brain/beeep"
	"github.com/jasonlvhit/gocron"
	"github.com/m7mdkamal/webwatcher/watcher"
)

func main() {
	watcher.InitDB()
	tasks := *watcher.SelectAllTasks()
	log.Printf("[webwatcher] found %d tasks", len(tasks))

	var wg sync.WaitGroup

	testfunc := func(task watcher.Task) {
		log.Printf("[%s] watcher started", task.Filter)
		results := task.Watcher.Run()
		log.Printf("[%s] watcher found %d", task.Filter, len(results))
		for _, result := range results {
			log.Printf("[%s] watcher found: %s", task.Filter, result.Title)

		}
		err := beeep.Notify("Title", "Message body", "assets/information.png")
		if err != nil {
			panic(err)
		}
		log.Printf("[%s] watcher ended. Should start after %d secs", task.Filter, task.Interval)
		// wg.Done()
	}

	for _, task := range tasks {
		fmt.Println(task)
		s := gocron.NewScheduler()
		s.Every(uint64(task.Interval)).Seconds().Do(testfunc, task)
		go s.Start()
		wg.Add(1)
	}
	wg.Wait()
}
