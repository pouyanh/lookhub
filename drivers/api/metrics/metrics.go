package metrics

import (
	"fmt"
	"net"

	"github.com/janstoon/toolbox/kareless"
	"github.com/janstoon/toolbox/tricks"

	"gitlab.snapp.ir/pouyanh/lookhub/settings"
)

var API kareless.DriverConstructor = apiDriver

func apiDriver(ss *kareless.Settings, ib *kareless.InstrumentBank, apps []kareless.Application) kareless.Driver {
	apiSs := settings.APIByName(ss, "metrics")

	return newServer(
		apiSs,
		tricks.Map(apiSs.Gateways, func(gwSs settings.Gateway) net.Listener {
			return kareless.ResolveInstrumentByType[net.Listener](ib, fmt.Sprintf("socket/%s", gwSs.Name))
		}),
	)
}
