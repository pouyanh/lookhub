package groper

import (
	"context"

	"gitlab.snapp.ir/pouyanh/lookhub/dnslv"
)

type domainRepository interface {
	GetDomain(ctx context.Context, name string) (*dnslv.Domain, error)
}
