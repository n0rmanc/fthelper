package plugins

import (
	"github.com/n0rmanc/fthelper/generator/v4/src/clusters"
	"github.com/n0rmanc/fthelper/shared/configs"
	"github.com/n0rmanc/fthelper/shared/fs"
	"github.com/n0rmanc/fthelper/shared/maps"
	"github.com/n0rmanc/fthelper/shared/runners"
)

func Json(data maps.Mapper, config maps.Mapper) *runners.Runner {
	return clusters.NewRunnerV2(data, config, func(p *clusters.ExecutorParameter) error {
		templates, err := fs.Build(fs.ToObject(p.Data.Zi("inputs"), p.Config), p.VarConfig)
		if err != nil {
			p.Logger.Error("cannot get inputs information")
			return err
		}

		output, err := fs.Build(fs.ToObject(p.Data.Zi("output"), p.Config), p.VarConfig)
		if err != nil {
			p.Logger.Error("cannot get output information")
			return err
		}

		content, err := configs.LoadConfigFromFileSystem(
			templates.Multiple(),
			p.Config,
			p.Data.Mi("merger"),
		)
		if err != nil {
			p.Logger.Error("cannot load template as json")
			return err
		}

		var outfile = output.Single()
		err = outfile.Build()
		if err != nil {
			p.Logger.Error("cannot build output file")
			return err
		}

		json, err := maps.ToFormatJson(content)
		if err != nil {
			p.Logger.Error("cannot create formatted json")
			return err
		}

		return outfile.Write(json)
	}, &clusters.Settings{
		DefaultWithCluster: true,
	})
}
