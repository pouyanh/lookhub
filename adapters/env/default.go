package env

import (
	"github.com/janstoon/toolbox/kareless"

	"gitlab.snapp.ir/pouyanh/lookhub"
)

func privileged(ss *kareless.Settings) []kareless.InstrumentCatalogue {
	return isolated(ss)
}

func isolated(_ *kareless.Settings) []kareless.InstrumentCatalogue {
	return []kareless.InstrumentCatalogue{
		{
			Names: []string{lookhub.ExternalServicesReady},
			Builder: func(ss *kareless.Settings, ib *kareless.InstrumentBank) kareless.Instrument {
				return true
			},
		},
	}
}
