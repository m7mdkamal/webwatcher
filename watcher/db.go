package watcher

import (
	"database/sql"
	"encoding/json"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteDataBase struct {
	*sql.DB
}

var db sqliteDataBase
var err error

func InitDB() {
	db.DB, err = sql.Open("sqlite3", "./webwatcher.db")
	checkErr(err)

}

func InsertTask(task *Task) int64 {
	stmt, err := db.Prepare("INSERT INTO tasks(name, watcherId, filter , parameters,interval, createdAt) values(?,?,?,?,?,?)")
	checkErr(err)

	parametersJson, _ := json.Marshal(task.Parameters)
	res, err := stmt.Exec(task.Name, task.WatcherID, task.Filter, parametersJson, task.Interval, time.Now().UTC())
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	return id
}

func SelectAllTasks() *[]Task {
	var tasks []Task
	rows, err := db.Query("SELECT name, watcherId, filter , parameters,interval, createdAt FROM tasks")
	defer rows.Close()
	checkErr(err)
	var (
		name           string
		watcherId      int
		filter         string
		parameters     []interface{}
		parametersJson string
		interval       int
		createdAt      int
	)

	for rows.Next() {
		rows.Scan(&name, &watcherId, &filter, &parametersJson, &interval, &createdAt)
		json.Unmarshal([]byte(parametersJson), &parameters)
		tasks = append(tasks, *NewTask(name, Reddit, filter, interval, parameters...))
	}
	return &tasks
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
