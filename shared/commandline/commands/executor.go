package commands

import (
	"github.com/n0rmanc/fthelper/shared/caches"
	"github.com/n0rmanc/fthelper/shared/commandline/models"
	"github.com/n0rmanc/fthelper/shared/loggers"
	"github.com/n0rmanc/fthelper/shared/maps"
)

type ExecutorParameter struct {
	Name   string
	Meta   *models.Metadata
	Config maps.Mapper
	Cache  *caches.Service
	Logger *loggers.Logger
	Args   []string
}

type Executor func(p *ExecutorParameter) error
