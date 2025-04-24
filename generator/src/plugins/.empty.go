package plugins

import (
	"github.com/n0rmanc/fthelper/generator/v4/src/clusters"
	"github.com/n0rmanc/fthelper/shared/maps"
	"github.com/n0rmanc/fthelper/shared/runners"
)

func Empty(data maps.Mapper, config maps.Mapper) *runners.Runner {
	return clusters.NewRunnerV2(data, config, func(p *clusters.ExecutorParameter) error {
		return nil
	}, &clusters.Settings{
		DefaultWithCluster: true,
	})
}
