package drivers

import (
	"github.com/janstoon/toolbox/tricks"

	"github.com/pouyanh/lookhub/drivers/api"
)

var Drivers = tricks.Flat(
	api.APIs,
)
