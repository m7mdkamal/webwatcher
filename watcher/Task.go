package watcher

type Task struct {
	Name       string
	WatcherID  WatcherType   `yaml:"watcherId"`
	Watcher    Watcher       `yaml:"watcher"`
	Filter     string        `yaml:"filter"`
	Interval   int           `yaml:"interval"`
	Parameters []interface{} `yaml:"parameters"`
}

func NewTask(name string, wt WatcherType, filter string, interval int, parameters ...interface{}) *Task {
	return &Task{
		Name:       name,
		WatcherID:  wt,
		Watcher:    NewWatcher(wt, filter, parameters...),
		Filter:     filter,
		Interval:   interval,
		Parameters: parameters,
	}
}
