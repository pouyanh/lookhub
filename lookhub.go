package lookhub

import (
	"github.com/janstoon/toolbox/bricks"
	"github.com/janstoon/toolbox/kareless"
)

// BundlesEssentials lets bundles to register their own Applications, Adapters and Drivers
// They become injected automatically
var BundlesEssentials = bricks.NewRegistry[[]kareless.Option]()

// ExternalServicesReady is a virtual dependency name which resolves immediately in non-testing operation modes.
// In testing operation mode it should get resolved whenever temporary external services are created and identified.
const ExternalServicesReady = "extern"
