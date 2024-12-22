package groper_test

import (
	"context"
	"testing"
	"time"

	"github.com/janstoon/toolbox/bricks"
	"github.com/janstoon/toolbox/kareless"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gitlab.snapp.ir/pouyanh/lookhub/dnslv"
	"gitlab.snapp.ir/pouyanh/lookhub/dnslv/groper"
)

func TestApplication_Lookup(t *testing.T) {
	var (
		app *groper.Application

		domains domainRepo
		dns     domainNameService
	)

	k := kareless.Compile().
		Equip(adapters...).
		Install(func(ss *kareless.Settings, ib *kareless.InstrumentBank) kareless.Application {
			app = groper.NewApp(ss, ib)

			return app
		})

	ctx, stop := run(
		context.Background(), k,
		func(ctx context.Context, ss *kareless.Settings, ib *kareless.InstrumentBank, apps []kareless.Application) error {
			domains = kareless.ResolveInstrumentByType[domainRepo](ib, "repo/dnslv/domain")
			dns = kareless.ResolveInstrumentByType[domainNameService](ib, "svc/dnslv/dns")

			return nil
		},
	)
	defer stop()

	{
		name := "pouyan.dev"
		dns[name] = []dnslv.ResourceRecord{
			{Type: "NS", Value: "ns1.example.com.", TTL: 21600},
			{Type: "NS", Value: "ns2.example.com.", TTL: 21600},
			{Type: "A", Value: "10.0.0.1", TTL: 14400},
			{Type: "MX", Value: "0 pouyan.dev.", TTL: 14400},
		}

		assert.NotContains(t, domains, name)
		domain, err := app.Lookup(ctx, name)
		require.NoError(t, err)
		assert.Contains(t, domains, name)
		assert.ElementsMatch(t,
			domain.Records(),
			[]dnslv.ResourceRecord{
				{Type: "NS", Value: "ns1.example.com.", TTL: 21600},
				{Type: "NS", Value: "ns2.example.com.", TTL: 21600},
				{Type: "A", Value: "10.0.0.1", TTL: 14400},
				{Type: "MX", Value: "0 pouyan.dev.", TTL: 14400},
			},
		)

		// use cache
		delete(dns, name)
		assert.NotContains(t, dns, name)
		domain, err = app.Lookup(ctx, name)
		require.NoError(t, err)
		assert.Contains(t, domains, name)
		assert.ElementsMatch(t,
			domain.Records(),
			[]dnslv.ResourceRecord{
				{Type: "NS", Value: "ns1.example.com.", TTL: 21600},
				{Type: "NS", Value: "ns2.example.com.", TTL: 21600},
				{Type: "A", Value: "10.0.0.1", TTL: 14400},
				{Type: "MX", Value: "0 pouyan.dev.", TTL: 14400},
			},
		)
	}
}

func run(ctx context.Context, k kareless.Kernel, hooks ...kareless.Hook) (context.Context, context.CancelFunc) {
	ctx, stop := context.WithCancel(ctx)
	started := make(chan bool)
	go func() {
		_ = k.
			AfterStart(
				func(ctx context.Context, ss *kareless.Settings, ib *kareless.InstrumentBank, apps []kareless.Application) error {
					for _, hook := range hooks {
						err := hook(ctx, ss, ib, apps)
						if err != nil {
							return err
						}
					}

					started <- true

					return nil
				},
			).Run(ctx)
	}()

	<-started

	return ctx, stop
}

var adapters = []kareless.InstrumentInjector{
	func(ss *kareless.Settings) []kareless.InstrumentCatalogue {
		return []kareless.InstrumentCatalogue{
			{
				Names: []string{"repo/dnslv/domain"},
				Builder: func(ss *kareless.Settings, ib *kareless.InstrumentBank) kareless.Instrument {
					return make(domainRepo)
				},
			},
			{
				Names: []string{"svc/dnslv/dns"},
				Builder: func(ss *kareless.Settings, ib *kareless.InstrumentBank) kareless.Instrument {
					return make(domainNameService)
				},
			},
		}
	},
}

type domainRepo map[string]domainDao

type domainDao struct {
	domain    *dnslv.Domain
	updatedAt time.Time
}

var _ groper.DomainRepository = (*domainRepo)(nil)

func (dd domainRepo) GetDomain(_ context.Context, name string, ttl time.Duration) (*dnslv.Domain, error) {
	v, ok := dd[name]
	if !ok || time.Now().After(v.updatedAt.Add(ttl)) {
		return nil, bricks.ErrNotFound
	}

	return v.domain, nil
}

func (dd domainRepo) SaveDomain(_ context.Context, domain *dnslv.Domain) error {
	dd[domain.Name().String()] = domainDao{
		domain:    domain,
		updatedAt: time.Now(),
	}

	return nil
}

type domainNameService map[string][]dnslv.ResourceRecord

var _ groper.DomainNameService = (*domainNameService)(nil)

func (svc domainNameService) QueryAllDNSRecords(_ context.Context, domainName string) ([]dnslv.ResourceRecord, error) {
	v, ok := svc[domainName]
	if !ok {
		return nil, bricks.ErrNotFound
	}

	return v, nil
}
