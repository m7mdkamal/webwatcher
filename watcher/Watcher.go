package watcher

import "github.com/m7mdkamal/webwatcher/model"

type Watcher interface {
	Run() []model.Result
}
