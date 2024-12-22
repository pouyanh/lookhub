package db

import (
	"fmt"

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

						return newDatabaseConnection(settings.DatabaseByName(ss, db.Name), settings.OperationMode(ss))
					},
				}
			},
		)
	},
}

func newDatabaseConnection(ss settings.Database, mode settings.Mode) any {
	return "TODO"
}
