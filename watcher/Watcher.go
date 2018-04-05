package watcher

import "github.com/m7mdkamal/webwatcher/model"

type WatcherWorker interface {
	Run() []model.Result
}
