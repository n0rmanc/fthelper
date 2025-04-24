package generators

import (
	"fmt"

	"github.com/n0rmanc/fthelper/generator/v4/src/plugins"
	"github.com/n0rmanc/fthelper/shared/maps"
	"github.com/n0rmanc/fthelper/shared/runners"
)

func GetRunner(data maps.Mapper, config maps.Mapper) (*runners.Runner, error) {
	var name = data.Si("type")
	switch name {
	case "bash":
		return plugins.Bash(data, config), nil
	case "json":
		return plugins.Json(data, config), nil
	case "create":
		return plugins.Create(data, config), nil
	case "copy":
		return plugins.Copy(data, config), nil
	case "template":
		return plugins.Template(data, config), nil
	case "strategy":
		return plugins.CStrategy(data, config), nil
	case "config":
		return plugins.CConfig(data, config), nil
	}

	return nil, fmt.Errorf("cannot found generator for type '%s'", name)
}
