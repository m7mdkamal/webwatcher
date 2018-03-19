package model

// Watcher is watcher info
type Watcher struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	CreatedAt int64  `db:"created_at"`
}

// Task is what we search for
type Task struct {
	ID         int64         `db:"id"`
	Name       string        `db:"name"`
	WatcherID  int64         `db:"watcher_id"`
	Filter     string        `db:"filter"`
	Interval   int           `db:"interval"`
	Parameters []interface{} `db:"parameters"`
}

// Result is record of information
type Result struct {
	ID      int64  `db:"id"`
	Title   string `db:"title"`
	Content string `db:"content"`
	URL     string `db:"url"`
	Time    int64  `db:"time"`
	TaskID  int64  `db:"task_id"`
}

// func NewTask(id int64, name string, wt WatcherType, filter string, interval int, parameters ...interface{}) *Task {
// 	return &Task{
// 		ID:         id,
// 		Name:       name,
// 		WatcherID:  wt,
// 		Watcher:    NewWatcher(wt, filter, parameters...),
// 		Filter:     filter,
// 		Interval:   interval,
// 		Parameters: parameters,
// 	}
// }
