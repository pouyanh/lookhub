package adapters

import (
	"github.com/janstoon/toolbox/tricks"

	"github.com/pouyanh/lookhub/adapters/db"
	"github.com/pouyanh/lookhub/adapters/env"
	"github.com/pouyanh/lookhub/adapters/socket"
)

var Adapters = tricks.Flat(
	env.Environment, db.Databases, socket.Socket,
)
