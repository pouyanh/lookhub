package adapters

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/janstoon/toolbox/kareless"

	"gitlab.snapp.ir/pouyanh/lookhub/dnslv/adapters/db"
	"gitlab.snapp.ir/pouyanh/lookhub/settings"
)

var Adapters = []kareless.InstrumentInjector{
	func(ss *kareless.Settings) []kareless.InstrumentCatalogue {
		return []kareless.InstrumentCatalogue{
			{
				Names: []string{"repo/dnslv/domain"},
				Builder: func(ss *kareless.Settings, ib *kareless.InstrumentBank) kareless.Instrument {
					return domainRepoAdapter(ss, ib)
				},
			},
			{
				Names: []string{"svc/dnslv/dns"},
				Builder: func(ss *kareless.Settings, ib *kareless.InstrumentBank) kareless.Instrument {
					return dnsClientAdapter(ss)
				},
			},
		}
	},
}

func domainRepoAdapter(ss *kareless.Settings, ib *kareless.InstrumentBank) domainRepo {
	conn := kareless.ResolveInstrumentByType[*pgxpool.Pool](ib, "db/dnslv")

	err := db.Migrate(conn)
	if err != nil {
		panic(err)
	}

	return newDomainRepo(conn)
}

func dnsClientAdapter(ss *kareless.Settings) dnsClient {
	return newDNSClient(settings.LookupServer(ss))
}
