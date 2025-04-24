package hooks

import "github.com/n0rmanc/fthelper/shared/maps"

// Hook action
type Hook func(config maps.Mapper) error
