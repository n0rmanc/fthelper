package freqtrade

import (
	"github.com/n0rmanc/fthelper/metric/v4/src/connection"
	"github.com/n0rmanc/fthelper/shared/datatype"
	"github.com/n0rmanc/fthelper/shared/maps"
)

const VERSION_CONST = "version"

func NewVersion() *version {
	return &version{}
}

type version struct{}

func (v *version) Name() string {
	return VERSION_CONST
}

func (p *version) Build(connector connection.Connector, connection *connection.Connection, history *datatype.Queue) (interface{}, error) {
	var target = make(maps.Mapper)
	err := connection.Http.GET(p.Name(), &target)
	if err != nil {
		return nil, err
	}

	return target.Se("version")
}

func ToVersion(connector connection.Connector) string {
	raw, err := connector.Connect(VERSION_CONST)
	if err != nil {
		return "v0.0.0"
	}
	return raw.(string)
}
