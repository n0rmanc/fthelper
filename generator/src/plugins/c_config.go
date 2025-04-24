package plugins

import (
	"strings"

	"github.com/n0rmanc/fthelper/generator/v4/src/clusters"
	"github.com/n0rmanc/fthelper/shared/configs"
	"github.com/n0rmanc/fthelper/shared/fs"
	"github.com/n0rmanc/fthelper/shared/maps"
	"github.com/n0rmanc/fthelper/shared/runners"
)

// TODO: support override config from environment variable
// CConfig is custom plugins for and only for freqtrade config
func CConfig(data maps.Mapper, config maps.Mapper) *runners.Runner {
	return clusters.NewRunnerV2(data, config, func(p *clusters.ExecutorParameter) error {
		input, err := fs.Build(fs.ToObject(p.Data.Zi("input"), p.Config), p.VarConfig)
		if err != nil {
			p.Logger.Error("cannot get input information")
			return err
		}

		var files = make([]fs.FileSystem, 0)
		if input.IsSingle() {
			directory, err := fs.NewDirectory(fs.Next(input.Single(), p.VarConfig.Si("config")))
			if err != nil {
				p.Logger.Error("cannot get find freqtrade configs template directory")
				return err
			}
			files = []fs.FileSystem{directory}
		} else if input.IsMultiple() {
			files = input.Multiple()
		}

		var config = configs.BuildClusterConfig(p.Data.Si("withEnv"), p.Config)
		content, err := configs.LoadConfigFromFileSystem(files, config, p.Data.Mi("merger"))
		if err != nil {
			p.Logger.Error("cannot load template data")
			return err
		}
		json, err := maps.ToFormatJson(content)
		if err != nil {
			p.Logger.Error("cannot format config to json")
			return err
		}

		var filename strings.Builder
		filename.WriteString("config")
		if p.Data.Si("suffix") != "" {
			filename.WriteString("-" + p.Data.Si("suffix"))
		}
		var cluster = p.Data.Si("cluster")
		if p.Data.Bo("clusterSuffix", false) && cluster != "" {
			filename.WriteString("-" + cluster)
		}
		filename.WriteString(".json")
		output, err := fs.Build(fs.ToObject(p.Data.Zi("output"), p.Config), p.VarConfig)
		if err != nil {
			p.Logger.Error("cannot get output information")
			return err
		}
		file, err := fs.NewFile(fs.Next(output.Single(), p.VarConfig.Si("userdata"), filename.String()))
		if err != nil {
			p.Logger.Error("cannot get find freqtrade configs directory")
			return err
		}

		err = file.Build()
		if err != nil {
			p.Logger.Error("cannot build output directory")
			return err
		}
		return file.Write(json)
	}, &clusters.Settings{
		DefaultWithCluster: true,
	})
}
