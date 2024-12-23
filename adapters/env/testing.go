package env

import (
	"fmt"

	"github.com/janstoon/toolbox/kareless"
	"github.com/janstoon/toolbox/kareless/std"
	"github.com/janstoon/toolbox/tricks"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/postgres"

	"gitlab.snapp.ir/pouyanh/lookhub"
	"gitlab.snapp.ir/pouyanh/lookhub/settings"
)

func testing() []kareless.InstrumentCatalogue {
	return []kareless.InstrumentCatalogue{
		{
			Names: []string{lookhub.ExternalServicesReady},
			Builder: func(ss *kareless.Settings, ib *kareless.InstrumentBank) kareless.Instrument {
				containers := tricks.Map(
					settings.Databases(ss),
					func(src settings.Database) *gnomock.Container {
						c, err := gnomock.Start(postgres.Preset(
							postgres.WithUser(src.Username, src.Password),
							postgres.WithDatabase(src.DbName),
						), gnomock.WithUseLocalImagesFirst())
						if err != nil {
							panic(err)
						}

						ss.Prepend(std.MapSettingSource{
							fmt.Sprintf("databases.%s.host", src.Name): c.Host,
							fmt.Sprintf("databases.%s.port", src.Name): c.DefaultPort(),
						})

						return c
					},
				)

				return containers
			},
		},
	}
}
