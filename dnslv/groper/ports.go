package groper

import (
	"context"
	"time"

	"github.com/pouyanh/lookhub/dnslv"
)

type domainRepository interface {
	GetUnexpiredDomain(ctx context.Context, name string, ttl time.Duration) (*dnslv.Domain, error)
	SyncDomain(ctx context.Context, domain *dnslv.Domain) error
}

type domainNameService interface {
	QueryAllDNSRecords(ctx context.Context, domainName string) ([]dnslv.ResourceRecord, error)
}
