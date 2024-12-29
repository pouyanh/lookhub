package groper

import (
	"context"
	"time"

	"github.com/janstoon/toolbox/kareless"

	"gitlab.snapp.ir/pouyanh/lookhub/dnslv"
	"gitlab.snapp.ir/pouyanh/lookhub/lutel"
	"gitlab.snapp.ir/pouyanh/lookhub/settings"
)

type Application struct {
	domains domainRepository
	dns     domainNameService

	cacheTTL time.Duration
}

func NewApp(ss *kareless.Settings, ib *kareless.InstrumentBank) *Application {
	return &Application{
		domains: kareless.ResolveInstrumentByType[domainRepository](ib, "repo/dnslv/domain"),
		dns:     kareless.ResolveInstrumentByType[domainNameService](ib, "svc/dnslv/dns"),

		cacheTTL: settings.LookupTTL(ss),
	}
}

func (app Application) Lookup(ctx context.Context, domainName string) (*dnslv.Domain, error) {
	domain, err := app.domains.GetUnexpiredDomain(ctx, domainName, app.cacheTTL)
	if err == nil {
		lutel.CacheHitCnt.Inc()
	} else {
		// todo: log the error and ignore
		lutel.CacheMissCnt.Inc()

		domain, err = dnslv.DomainByName(dnslv.DomainName(domainName))
		if err != nil {
			return nil, err
		}
	}

	if len(domain.Records()) > 0 {
		return domain, nil
	}

	ctxReq, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	records, err := app.dns.QueryAllDNSRecords(ctxReq, domainName)
	if err != nil {
		return nil, err
	}

	err = domain.AddRecords(records...)
	if err != nil {
		return nil, err
	}

	err = app.domains.SyncDomain(ctx, domain)
	if err != nil {
		// todo: log the error and ignore
	}

	return domain, nil
}
