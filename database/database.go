package database

import "github.com/m7mdkamal/webwatcher/model"

// Database is interface to CRUD data in database
type Database interface {

	// GetWatcherByTask get watcher by task id
	GetWatcherByTask(task *model.Task) (*model.Watcher, error)

	// CreateTask creates new task in db
	CreateTask(task *model.Task) (int64, error)

	// GetTasks get tasks from db
	GetTasks() ([]model.Task, error)

	// DeleteTask delete task from database
	DeleteTask(task *model.Task) error

	// CreateResult creates new result
	CreateResult(result *model.Result) (int64, error)

	// GetResultsByTask fetch results of a specific task
	GetResultsByTask(task *model.Task) ([]model.Result, error)
}
