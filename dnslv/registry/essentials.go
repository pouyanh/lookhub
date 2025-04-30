package registry

import (
	"github.com/janstoon/toolbox/kareless"

	"github.com/pouyanh/lookhub"
	"github.com/pouyanh/lookhub/dnslv/adapters"
	"github.com/pouyanh/lookhub/dnslv/groper"
)

func init() {
	err := lookhub.BundlesEssentials.Register("dnslv", []kareless.Option{
		kareless.Equipment(adapters.Adapters...),

		kareless.Installer(func(ss *kareless.Settings, ib *kareless.InstrumentBank) kareless.Application {
			return groper.NewApp(ss, ib)
		}),
	})
	if err != nil {
		panic(err)
	}
}
