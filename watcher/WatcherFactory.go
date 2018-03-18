package watcher

type WatcherType int

const (
	Reddit WatcherType = iota + 1
)

func IntToWatcherType(i int) WatcherType {
	switch i {
	case 1:
		return Reddit
	}
	return Reddit
}

func NewWatcher(wt WatcherType, filter string, parameters ...interface{}) Watcher {
	switch wt {
	case Reddit:
		filter := filter
		subReddit := parameters[0].(string)
		return NewRedditWatcher(subReddit, filter)
	}
	return nil
}
