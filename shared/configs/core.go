package configs

import (
	"github.com/n0rmanc/fthelper/shared/loggers"
	"github.com/n0rmanc/fthelper/shared/maps"
)

func New(name string, config maps.Mapper) *Builder {
	return &Builder{
		name:     name,
		config:   config,
		override: maps.New(),
		strategy: maps.New(),

		logger: loggers.Get("config", "builder"),
	}
}
