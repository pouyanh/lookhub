package socket

import (
	"errors"
	"fmt"
	"net"

	"github.com/janstoon/toolbox/bricks"
	"github.com/janstoon/toolbox/kareless"
	"github.com/janstoon/toolbox/tricks"

	"gitlab.snapp.ir/pouyanh/lookhub/settings"
)

var Socket = []kareless.InstrumentInjector{
	func(ss *kareless.Settings) []kareless.InstrumentCatalogue {
		return tricks.Map(settings.Gateways(ss), func(gw settings.Gateway) kareless.InstrumentCatalogue {
			return kareless.InstrumentCatalogue{
				Names: []string{
					fmt.Sprintf("socket/%s", gw.Name),
				},
				Builder: func(ss *kareless.Settings, ib *kareless.InstrumentBank) kareless.Instrument {
					return newSocket(gw)
				},
			}
		})
	},
}

func newSocket(gw settings.Gateway) net.Listener {
	l, err := net.Listen(gw.Proto, gw.Address)
	if err != nil {
		panic(errors.Join(bricks.ErrUnavailable, err))
	}

	return l
}
