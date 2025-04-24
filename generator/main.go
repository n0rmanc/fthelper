package main

import (
	"os"

	"github.com/n0rmanc/fthelper/generator/v4/src/generators"
	"github.com/n0rmanc/fthelper/shared/caches"
	"github.com/n0rmanc/fthelper/shared/commandline"
	"github.com/n0rmanc/fthelper/shared/commandline/commands"
	"github.com/n0rmanc/fthelper/shared/commandline/models"
	"github.com/n0rmanc/fthelper/shared/commandline/plugins"
)

var (
	name    string = "ftgenerator"
	version string = "dev"
	commit  string = "none"
	date    string = "unknown"
	builtBy string = "manually"
)

func main() {
	var cmd = commandline.New(caches.Global, &models.Metadata{
		Name:    name,
		Version: version,
		Commit:  commit,
		Date:    date,
		BuiltBy: builtBy,
	}).
		Plugin(plugins.SupportVersion).
		Plugin(plugins.SupportFSVar).
		Plugin(plugins.SupportVar).
		Plugin(plugins.SupportDotEnv). // dotenv must come before config
		Plugin(plugins.SupportListConfig).
		Plugin(plugins.SupportConfig).
		Plugin(plugins.SupportCluster).  // cluster must come after config
		Plugin(plugins.SupportLogLevel). // log-level must come after config
		Plugin(plugins.SupportBanner).
		Command(&commands.Command{
			Name: commands.DEFAULT,
			Executor: func(p *commands.ExecutorParameter) error {
				return generators.New(p.Cache, p.Config).Start()
			},
		})

	var err = cmd.Start(os.Args)
	if err != nil {
		panic(err)
	}
}
