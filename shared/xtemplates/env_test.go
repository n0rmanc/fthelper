package xtemplates_test

import (
	"testing"

	"github.com/n0rmanc/fthelper/shared/maps"
	"github.com/n0rmanc/fthelper/shared/xtemplates"
	"github.com/n0rmanc/fthelper/shared/xtests"
)

func TestByCluster(t *testing.T) {
	var assertion = xtests.New(t)
	var config = maps.New().
		Set("base", "default").
		Set("_.1A.base", "1st ant")

	assertion.NewName("override occurred").
		WithExpected("1st ant").
		WithActualAndError(xtemplates.Text("{{ byCluster . `1A` `base` }}", config)).
		MustEqual()

	assertion.NewName("fallback to default").
		WithExpected("default").
		WithActualAndError(xtemplates.Text("{{ byCluster . `3A` `base` }}", config)).
		MustEqual()
}
