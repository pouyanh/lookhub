package api

import (
	"github.com/janstoon/toolbox/kareless"

	"gitlab.snapp.ir/pouyanh/lookhub/drivers/api/metrics"
	"gitlab.snapp.ir/pouyanh/lookhub/drivers/api/user"
)

var APIs = []kareless.DriverConstructor{
	user.API,
	metrics.API,
}
