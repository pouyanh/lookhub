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
	)

	k := kareless.Compile().
		Equip(adapters...).
		Equip(func(ss *kareless.Settings) []kareless.InstrumentCatalogue {
			return []kareless.InstrumentCatalogue{
				{
					Names: []string{"_test"},
					Builder: func(ss *kareless.Settings, ib *kareless.InstrumentBank) kareless.Instrument {
						domains = kareless.ResolveInstrumentByType[domainRepo](ib, "repo/dnslv/domain")

						return true
					},
				},
			}
		}).
		Install(func(ss *kareless.Settings, ib *kareless.InstrumentBank) kareless.Application {
			app = groper.NewApp(ss, ib)

			return app
		})

	ctx, stop := run(context.Background(), k)
	defer stop()

	{
		assert.NotContains(t, domains, "pouyan.dev")
		domain, err := app.Lookup(ctx, "pouyan.dev")
		require.NoError(t, err)
		assert.Contains(t, domains, "pouyan.dev")
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

func run(ctx context.Context, k kareless.Kernel) (context.Context, context.CancelFunc) {
	ctx, stop := context.WithCancel(ctx)
	started := make(chan bool)
	go func() {
		_ = k.
			AfterStart(
				func(_ context.Context, _ *kareless.Settings, _ *kareless.InstrumentBank, _ []kareless.Application) error {
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

func (dd domainRepo) SaveDomain(ctx context.Context, domain *dnslv.Domain) error {
	dd[domain.Name().String()] = domainDao{
		domain:    domain,
		updatedAt: time.Now(),
	}

	return nil
}
