package user

import (
	"context"

	"gitlab.snapp.ir/pouyanh/lookhub/dnslv"
)

type DNSLVLookup interface {
	Lookup(ctx context.Context, domainName string) (*dnslv.Domain, error)
}
