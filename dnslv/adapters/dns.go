package adapters

import (
	"context"

	"github.com/janstoon/toolbox/bricks"
	"github.com/janstoon/toolbox/kareless"
	"github.com/miekg/dns"

	"gitlab.snapp.ir/pouyanh/lookhub/dnslv"
)

type dnsClient struct {
	base *dns.Client
}

func newDNSClient(ss *kareless.Settings, ib *kareless.InstrumentBank) dnsClient {
	c := dnsClient{
		base: new(dns.Client),
	}

	return c
}

func (c dnsClient) QueryAllDNSRecords(ctx context.Context, name string) ([]dnslv.ResourceRecord, error) {
	return nil, bricks.ErrUnimplemented
}
