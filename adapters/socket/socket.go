package socket

import (
	"errors"
	"fmt"
	"net"

	"github.com/janstoon/toolbox/bricks"
	"github.com/janstoon/toolbox/kareless"
	"github.com/janstoon/toolbox/tricks"

	"github.com/pouyanh/lookhub/settings"
)

var Socket = []kareless.InstrumentInjector{
	func(ss *kareless.Settings) []kareless.InstrumentCatalogue {
		return tricks.Map(settings.Gateways(ss), func(gwSs settings.Gateway) kareless.InstrumentCatalogue {
			return kareless.InstrumentCatalogue{
				Names: []string{
					fmt.Sprintf("socket/%s", gwSs.Name),
				},
				Builder: func(ss *kareless.Settings, ib *kareless.InstrumentBank) kareless.Instrument {
					return newSocket(gwSs)
				},
			}
		})
	},
}

func newSocket(gwSs settings.Gateway) net.Listener {
	l, err := net.Listen(gwSs.Proto, gwSs.Address)
	if err != nil {
		panic(errors.Join(bricks.ErrUnavailable, err))
	}

	return l
}
