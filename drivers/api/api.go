package api

import (
	"github.com/janstoon/toolbox/kareless"

	"github.com/pouyanh/lookhub/drivers/api/metrics"
	"github.com/pouyanh/lookhub/drivers/api/user"
)

var APIs = []kareless.DriverConstructor{
	user.API,
	metrics.API,
}
