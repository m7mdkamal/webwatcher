package watcher

const (
	Reddit int64 = iota + 1
)

func IntToWatcherType(i int64) int64 {
	switch i {
	case 1:
		return Reddit
	}
	return Reddit
}

func NewWatcher(wt int64, filter string, parameters ...interface{}) WatcherWorker {
	switch wt {
	case Reddit:
		filter := filter
		subReddit := parameters[0].(string)
		return NewRedditWatcher(subReddit, filter)
	}
	return nil
}
