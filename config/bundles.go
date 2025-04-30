package config

import (
	"maps"
	"slices"

	"github.com/janstoon/toolbox/kareless"

	"github.com/pouyanh/lookhub"
)

// BundlesEssentials returns registered dependencies of allowing bundles
func BundlesEssentials() []kareless.Option {
	return slices.Concat(slices.Collect(maps.Values(lookhub.BundlesEssentials.All()))...)
}
