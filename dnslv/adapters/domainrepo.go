package adapters

import (
	"context"
	"time"

	"github.com/janstoon/toolbox/bricks"
	"github.com/janstoon/toolbox/kareless"
	"github.com/janstoon/toolbox/tricks"

	"gitlab.snapp.ir/pouyanh/lookhub/dnslv"
)

type repoDomain struct {
	// todo: db connection
}

func newDomainRepo(ss *kareless.Settings, ib *kareless.InstrumentBank) repoDomain {
	return repoDomain{
		// todo: db connection
	}
}

func (repo repoDomain) GetDomain(ctx context.Context, name string, ttl time.Duration) (*dnslv.Domain, error) {
	var dst dnslv.Domain

	return tricks.ValPtr(dst), bricks.ErrUnimplemented
}

func (repo repoDomain) SaveDomain(ctx context.Context, domain dnslv.Domain) error {
	return bricks.ErrUnimplemented
}
