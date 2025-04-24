package clients

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/n0rmanc/fthelper/shared/maps"
)

type Query struct {
	data maps.Mapper
}

func (c *Query) Get(key string) url.Values {
	var values = make(url.Values)

	var mapper = c.data.Mi(key)
	mapper.ForEach(func(key string, value interface{}) {
		values.Set(key, fmt.Sprintf("%v", value))
	})

	return values
}

func (c *Query) Json() string {
	var j, err = json.Marshal(c.data)
	if err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}

func newQuery(query maps.Mapper) *Query {
	return &Query{
		data: query,
	}
}
