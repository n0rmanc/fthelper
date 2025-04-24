package commandline

import (
	"github.com/n0rmanc/fthelper/shared/caches"
	"github.com/n0rmanc/fthelper/shared/commandline/commands"
	"github.com/n0rmanc/fthelper/shared/commandline/flags"
	"github.com/n0rmanc/fthelper/shared/commandline/hooks"
	"github.com/n0rmanc/fthelper/shared/commandline/models"
	"github.com/n0rmanc/fthelper/shared/commandline/plugins"
	"github.com/n0rmanc/fthelper/shared/loggers"
)

func New(cache *caches.Service, metadata *models.Metadata) *cli {
	return &cli{
		Metadata: metadata,
		flags:    flags.New(),
		commands: commands.New(),
		hooks:    hooks.New(),
		plugins:  plugins.New(),

		cache:  cache,
		logger: loggers.Get("commandline"),
	}
}
