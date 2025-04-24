package clusters

import (
	"github.com/n0rmanc/fthelper/shared/loggers"
	"github.com/n0rmanc/fthelper/shared/maps"
)

type ExecutorParameter struct {
	// Executor index, this always unique
	Index int

	// Name of the executor
	Name string

	// type of current executor
	Type string

	// whole configuration mapper
	Config maps.Mapper

	// generator data
	Data maps.Mapper

	// fs configuration
	VarConfig maps.Mapper

	// helper for logging message
	Logger *loggers.Logger
}

type Executor func(p *ExecutorParameter) error
