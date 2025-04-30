package user

import (
	"context"

	"github.com/pouyanh/lookhub/dnslv"
)

type DNSLVLookup interface {
	Lookup(ctx context.Context, domainName string) (*dnslv.Domain, error)
}
