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
	dns     domainNameService
}

func NewApp(ss *kareless.Settings, ib *kareless.InstrumentBank) *Application {
	return &Application{
		domains: kareless.ResolveInstrumentByType[domainRepository](ib, "repo/dnslv/domain"),
		dns:     kareless.ResolveInstrumentByType[domainNameService](ib, "svc/dnslv/dns"),
	}
}

func (app Application) Lookup(ctx context.Context, name string) (*dnslv.Domain, error) {
	domain, err := app.domains.GetDomain(ctx, name, domainTtl)
	if err == nil {
		// todo: increase cache hit metric
	} else {
		// todo: log the error and ignore
		// todo: increase cache miss metric

		domain, err = dnslv.DomainByName(dnslv.DomainName(name))
		if err != nil {
			return nil, err
		}
	}

	if len(domain.Records()) > 0 {
		return domain, nil
	}

	ctxReq, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	records, err := app.dns.QueryAllDNSRecords(ctxReq, name)
	if err != nil {
		return nil, err
	}

	err = domain.AddRecords(records...)
	if err != nil {
		return nil, err
	}

	err = app.domains.SaveDomain(ctx, domain)
	if err != nil {
		// todo: log the error and ignore
	}

	return domain, nil
}
