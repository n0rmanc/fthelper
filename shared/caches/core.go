package caches

import (
	"sync"

	"github.com/n0rmanc/fthelper/shared/loggers"
)

func New() *Service {
	return &Service{
		caches: make(map[string]*Data),
		mutex:  sync.RWMutex{},
		logger: loggers.Get("cache", "service"),
	}
}

var Global = New()
