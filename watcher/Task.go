package watcher

type Task struct {
	ID         int64
	Name       string
	WatcherID  WatcherType   `yaml:"watcherId"`
	Watcher    Watcher       `yaml:"watcher"`
	Filter     string        `yaml:"filter"`
	Interval   int           `yaml:"interval"`
	Parameters []interface{} `yaml:"parameters"`
}

func NewTask(id int64, name string, wt WatcherType, filter string, interval int, parameters ...interface{}) *Task {
	return &Task{
		ID:         id,
		Name:       name,
		WatcherID:  wt,
		Watcher:    NewWatcher(wt, filter, parameters...),
		Filter:     filter,
		Interval:   interval,
		Parameters: parameters,
	}
}
