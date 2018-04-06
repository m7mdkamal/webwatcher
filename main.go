package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/m7mdkamal/webwatcher/cmd/webwatcher"
	"github.com/m7mdkamal/webwatcher/database"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, err := database.InitSQLiteDatabase(os.Getenv("SQLITE_DATABASE_PATH"))
	if err != nil {
		panic(err)
	}

	cmd.DB = db

	cmd.Execute()
}

// func main2() {

// 	// tasks, _ := db.GetTasks()
// 	// log.Printf("[webwatcher] found %d tasks", len(tasks))

// 	var wg sync.WaitGroup

// 	testfunc := func(task model.Task) {
// 		log.Printf("[%s] watcher started", task.Name)
// 		results := task.Watcher.Run()
// 		log.Printf("[%s] watcher found %d", task.Name, len(results))
// 		for _, result := range results {
// 			log.Printf("[%s] watcher found: %s", task.Name, result.Title)
// 			result.TaskId = task.ID
// 			db.CreateResult(&result)
// 		}
// 		err := beeep.Notify("Title", "Message body", "assets/information.png")
// 		if err != nil {
// 			panic(err)
// 		}
// 		log.Printf("[%s] watcher ended. Should start after %d secs", task.Name, task.Interval)
// 		// wg.Done()
// 	}

// 	for _, task := range tasks {
// 		s := gocron.NewScheduler()
// 		s.Every(uint64(task.Interval)).Seconds().Do(testfunc, task)
// 		go s.Start()
// 		wg.Add(1)
// 	}

// 	wg.Wait()
// }
