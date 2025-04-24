package plugins

import (
	"github.com/n0rmanc/fthelper/shared/commandline/commands"
	"github.com/n0rmanc/fthelper/shared/commandline/flags"
	"github.com/n0rmanc/fthelper/shared/commandline/hooks"
	"github.com/n0rmanc/fthelper/shared/commandline/models"
	"github.com/n0rmanc/fthelper/shared/loggers"
	"github.com/n0rmanc/fthelper/shared/maps"
)

type PluginParameter struct {
	Metadata models.Metadata

	NewCommand commands.Creator
	NewFlags   flags.Creator
	NewHook    hooks.Creator

	Config maps.Mapper
	Logger *loggers.Logger
}

type Plugin func(p *PluginParameter) error
