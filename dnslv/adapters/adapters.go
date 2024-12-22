package adapters

import "github.com/janstoon/toolbox/kareless"

var Adapters = []kareless.InstrumentInjector{
	func(ss *kareless.Settings) []kareless.InstrumentCatalogue {
		return []kareless.InstrumentCatalogue{
			{
				Names: []string{"repo/dnslv/domain"},
				Builder: func(ss *kareless.Settings, ib *kareless.InstrumentBank) kareless.Instrument {
					return newDomainRepo(ss, ib)
				},
			},
		}
	},
}
