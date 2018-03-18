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
	rows, err := db.Query("SELECT id, name, watcherId, filter , parameters,interval, createdAt FROM tasks")
	defer rows.Close()
	checkErr(err)
	var (
		id             int64
		name           string
		watcherId      int
		filter         string
		parameters     []interface{}
		parametersJson string
		interval       int
		createdAt      int
	)

	for rows.Next() {
		rows.Scan(&id, &name, &watcherId, &filter, &parametersJson, &interval, &createdAt)
		json.Unmarshal([]byte(parametersJson), &parameters)
		tasks = append(tasks, *NewTask(id, name, Reddit, filter, interval, parameters...))
	}
	return &tasks
}

func InsertResult(result *Result) int64 {
	stmt, err := db.Prepare("INSERT INTO results(taskId, title, content , url,time) values(?,?,?,?,?)")
	checkErr(err)

	res, err := stmt.Exec(result.TaskId, result.Title, result.Content, result.Url, result.Time)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	return id
}

func GetLastResultByTask(task *Task) *Result {
	rows, err := db.Query("SELECT title, content , url,time FROM results where taskId = ? order by id desc limit 1", task.ID)
	defer rows.Close()
	checkErr(err)

	var (
		result Result
	)

	for rows.Next() {
		rows.Scan(&result.Title, &result.Content, &result.Url, &result.Time)
	}
	return &result
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
