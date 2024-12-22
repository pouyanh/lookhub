package groper

import (
	"context"
	"time"

	"github.com/janstoon/toolbox/kareless"

	"gitlab.snapp.ir/pouyanh/lookhub/dnslv"
)

const domainTtl = 5 * time.Minute

type Application struct {
	domains domainRepository
}

func NewApp(ss *kareless.Settings, ib *kareless.InstrumentBank) *Application {
	return &Application{
		domains: kareless.ResolveInstrumentByType[domainRepository](ib, "repo/dnslv/domain"),
	}
}

func (app Application) Lookup(ctx context.Context, name string) (*dnslv.Domain, error) {
	domain, err := app.domains.GetDomain(ctx, name, domainTtl)
	if err != nil {
		// todo: log the error and ignore

		domain, err = dnslv.DomainByName(dnslv.DomainName(name))
		if err != nil {
			return nil, err
		}
	}

	if len(domain.Records()) == 0 {
		// todo: fetch all
	}

	err = app.domains.SaveDomain(ctx, domain)
	if err != nil {
		// todo: log the error and ignore
	}

	return domain, nil
}
