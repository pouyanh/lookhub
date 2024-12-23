package env

import (
	"github.com/janstoon/toolbox/kareless"

	"gitlab.snapp.ir/pouyanh/lookhub/settings"
)

var Environment = []kareless.InstrumentInjector{
	func(ss *kareless.Settings) []kareless.InstrumentCatalogue {
		switch settings.OperationMode(ss) {
		case settings.Testing:
			return testing()

		case settings.Privileged:
			return privileged()
		}

		return isolated()
	},
}
