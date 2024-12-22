package adapters

import (
	"github.com/janstoon/toolbox/tricks"

	"gitlab.snapp.ir/pouyanh/lookhub/adapters/db"
	"gitlab.snapp.ir/pouyanh/lookhub/adapters/env"
	"gitlab.snapp.ir/pouyanh/lookhub/adapters/socket"
)

var Adapters = tricks.Flat(
	env.Environment, db.Databases, socket.Socket,
)
