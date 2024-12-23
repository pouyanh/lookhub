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
			func(dbSs settings.Database) kareless.InstrumentCatalogue {
				return kareless.InstrumentCatalogue{
					Names: []string{
						fmt.Sprintf("db/%s", dbSs.Name),
					},
					Builder: func(ss *kareless.Settings, ib *kareless.InstrumentBank) kareless.Instrument {
						ib.Resolve(lookhub.ExternalServicesReady, func(v any) bool { return true })

						return dbConnAdapter(settings.DatabaseByName(ss, dbSs.Name), settings.OperationMode(ss))
					},
				}
			},
		)
	},
}

func dbConnAdapter(dbSs settings.Database, modeSs settings.Mode) any {
	conn, err := newDBConnection(dbSs, modeSs)
	if err != nil {
		panic(err)
	}

	return conn
}

func newDBConnection(dbSs settings.Database, modeSs settings.Mode) (any, error) {
	switch dbSs.Adapter {
	case "postgresql", "postgres", "pg":
		return newPgxConnectionPool(dbSs, modeSs)
	}

	return nil, bricks.ErrUnimplemented
}
