package groper

import (
	"context"
	"time"

	"gitlab.snapp.ir/pouyanh/lookhub/dnslv"
)

type domainRepository interface {
	GetDomain(ctx context.Context, name string, ttl time.Duration) (*dnslv.Domain, error)
	SaveDomain(ctx context.Context, domain *dnslv.Domain) error
}

type domainNameService interface {
	QueryAllDNSRecords(ctx context.Context, domainName string) ([]dnslv.ResourceRecord, error)
}
