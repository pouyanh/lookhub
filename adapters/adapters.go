package adapters

import (
	"github.com/janstoon/toolbox/tricks"

	"gitlab.snapp.ir/pouyanh/lookhub/adapters/db"
	"gitlab.snapp.ir/pouyanh/lookhub/adapters/env"
)

var Adapters = tricks.Flat(
	env.Environment, db.Databases,
)
