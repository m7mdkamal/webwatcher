package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/m7mdkamal/webwatcher/model"
	_ "github.com/mattn/go-sqlite3"
)

// SQLiteDatabase is the implementation of Database interface for sqlite3 database
type SQLiteDatabase struct {
	sqlx.DB
}

// InitSQLiteDatabase for create and connect to the database file
func InitSQLiteDatabase(databasePath string) (*SQLiteDatabase, error) {
	var err error
	// connect
	db := sqlx.MustConnect("sqlite3", databasePath)
	tx := db.MustBegin()

	// if panic happen during the creation rollback all the changes
	defer func() {
		if r := recover(); r != nil {
			panicErr, _ := r.(error)
			fmt.Println("Database error: ", panicErr)
			tx.Rollback()
			db = nil
			err = panicErr
		}
	}()

	// watcher table
	tx.MustExec(`CREATE TABLE IF NOT EXISTS watchers (
		id	INTEGER PRIMARY KEY AUTOINCREMENT,
		name	TEXT,
		created_at	INTEGER
	);`)

	// tasks table
	tx.MustExec(`CREATE TABLE IF NOT EXISTS tasks (
		id	INTEGER PRIMARY KEY AUTOINCREMENT,
		name	TEXT,
		watcher_id	INTEGER NOT NULL,
		filter	TEXT,
		parameters	TEXT,
		created_at	INTEGER,
		interval	INTEGER DEFAULT 60
	);`)

	// results table
	tx.MustExec(`CREATE TABLE IF NOT EXISTS results (
		id	INTEGER PRIMARY KEY AUTOINCREMENT,
		task_id	INTEGER NOT NULL,
		title	TEXT,
		content	TEXT,
		url	TEXT,
		time	INTEGER
	);`)

	// commit the changes
	err = tx.Commit()
	checkError(err)

	return &SQLiteDatabase{*db}, err
}

// GetWatcherByTask get watcher by task id
func (db *SQLiteDatabase) GetWatcherByTask(task *model.Task) (*model.Watcher, error) {
	watchers := []model.Watcher{}
	stmtGet, err := db.Preparex("SELECT id , name , created_at FROM watchers where id = ? limit 1")
	if err != nil {
		return nil, err
	}
	defer stmtGet.Close()
	err = stmtGet.Select(&watchers, task.ID)

	// exclude no rows error
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	fmt.Printf("%s\n", watchers)

	return &watchers[0], nil
}

// CreateTask creates new task in db
func (db *SQLiteDatabase) CreateTask(task *model.Task) (taskID int64, err error) {
	// Check name and interval
	if task.Name == "" {
		return -1, fmt.Errorf("task name must not be empty")
	}

	if task.Interval == 0 {
		return -1, fmt.Errorf("interval is required")
	}

	// Prepare transaction
	tx, err := db.Beginx()
	if err != nil {
		return -1, err
	}

	// Make sure to rollback if panic ever happened
	defer func() {
		if r := recover(); r != nil {
			panicErr, _ := r.(error)
			tx.Rollback()

			taskID = -1
			err = panicErr
		}
	}()

	// encode parameters to json
	parametersJson, _ := json.Marshal(task.Parameters)

	// Save article to database
	res := tx.MustExec(`INSERT INTO tasks
		(name, watcher_id, filter , parameters,
		interval, created_at) values(?,?,?,?,?,?)`,
		task.Name, task.WatcherID, task.Filter, parametersJson, task.Interval, time.Now().UTC())

	// Get last inserted ID
	taskID, err = res.LastInsertId()
	checkError(err)

	// Commit transaction
	err = tx.Commit()
	checkError(err)

	return taskID, err
}

// GetTasks get tasks from db
func (db *SQLiteDatabase) GetTasks() ([]model.Task, error) {

	tasks := []model.Task{}
	query := "SELECT id, name, watcher_id, filter, parameters, interval, created_at FROM tasks"

	err := db.Select(&tasks, query)
	// exclude no rows error
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return tasks, nil
}

// DeleteTask delete task from database
func (db *SQLiteDatabase) DeleteTask(task *model.Task) error {
	panic("Not implemented")
}

// DeleteTasks delete all tasks from database
func (db *SQLiteDatabase) DeleteTasks() error {
	panic("Not implemented")
}

// CreateResult creates new result
func (db *SQLiteDatabase) CreateResult(result *model.Result) (resultID int64, err error) {
	// Prepare transaction
	tx, err := db.Beginx()
	if err != nil {
		return -1, err
	}

	// Make sure to rollback if panic ever happened
	defer func() {
		if r := recover(); r != nil {
			panicErr, _ := r.(error)
			tx.Rollback()

			resultID = -1
			err = panicErr
		}
	}()

	// Save article to database
	res := tx.MustExec(`INSERT INTO results(task_id, title, content , url,time) values(?,?,?,?,?)`,
		result.TaskID, result.Title, result.Content, result.URL, result.Time)

	// Get last inserted ID
	resultID, err = res.LastInsertId()
	checkError(err)

	// Commit transaction
	err = tx.Commit()
	checkError(err)

	// return resultID, err
	return
}

// GetResultsByTask fetch results of a specific task
func (db *SQLiteDatabase) GetResultsByTask(task *model.Task) ([]model.Result, error) {

	results := []model.Result{}
	stmtGetResults, err := db.Preparex("SELECT title, content, url, time FROM results where taskId = ? order by id desc limit 1")
	if err != nil {
		return nil, err
	}
	defer stmtGetResults.Close()
	err = stmtGetResults.Select(&results, task.ID)

	// exclude no rows error
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return results, nil
}

// DeleteResults delete all results from database
func (db *SQLiteDatabase) DeleteResults() error {
	panic("Not implemented yet")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
