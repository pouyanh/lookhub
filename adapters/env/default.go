package env

import (
	"github.com/janstoon/toolbox/kareless"

	"github.com/pouyanh/lookhub"
)

func privileged() []kareless.InstrumentCatalogue {
	return isolated()
}

func isolated() []kareless.InstrumentCatalogue {
	return []kareless.InstrumentCatalogue{
		{
			Names: []string{lookhub.ExternalServicesReady},
			Builder: func(ss *kareless.Settings, ib *kareless.InstrumentBank) kareless.Instrument {
				return true
			},
		},
	}
}
