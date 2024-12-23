package db

import (
	"fmt"

	"github.com/janstoon/toolbox/bricks"
	"github.com/janstoon/toolbox/kareless"
	"github.com/janstoon/toolbox/tricks"

	"gitlab.snapp.ir/pouyanh/lookhub"
	"gitlab.snapp.ir/pouyanh/lookhub/settings"
)

var Databases = []kareless.InstrumentInjector{
	func(ss *kareless.Settings) []kareless.InstrumentCatalogue {
		return tricks.Map(
			settings.Databases(ss),
			func(db settings.Database) kareless.InstrumentCatalogue {
				return kareless.InstrumentCatalogue{
					Names: []string{
						fmt.Sprintf("db/%s", db.Name),
					},
					Builder: func(ss *kareless.Settings, ib *kareless.InstrumentBank) kareless.Instrument {
						ib.Resolve(lookhub.ExternalServicesReady, func(v any) bool { return true })

						return dbConnAdapter(settings.DatabaseByName(ss, db.Name), settings.OperationMode(ss))
					},
				}
			},
		)
	},
}

func dbConnAdapter(db settings.Database, mode settings.Mode) any {
	conn, err := newDBConnection(db, mode)
	if err != nil {
		panic(err)
	}

	return conn
}

func newDBConnection(db settings.Database, mode settings.Mode) (any, error) {
	switch db.Adapter {
	case "postgresql", "postgres", "pg":
		return newPgxConnectionPool(db, mode)
	}

	return nil, bricks.ErrUnimplemented
}
