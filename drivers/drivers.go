package drivers

import (
	"github.com/janstoon/toolbox/tricks"

	"gitlab.snapp.ir/pouyanh/lookhub/drivers/api"
)

var Drivers = tricks.Flat(
	api.APIs,
)
