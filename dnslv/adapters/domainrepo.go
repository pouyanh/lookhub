package adapters

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/janstoon/toolbox/bricks"
	"github.com/janstoon/toolbox/tricks"

	"gitlab.snapp.ir/pouyanh/lookhub/dnslv"
)

type domainRepo struct {
	conn *pgxpool.Pool
}

func newDomainRepo(conn *pgxpool.Pool) domainRepo {
	return domainRepo{
		conn: conn,
	}
}

func (repo domainRepo) GetUnexpiredDomain(ctx context.Context, name string, ttl time.Duration) (*dnslv.Domain, error) {
	var dst dnslv.Domain

	return tricks.ValPtr(dst), bricks.ErrUnimplemented
}

func (repo domainRepo) SyncDomain(ctx context.Context, domain *dnslv.Domain) error {
	return bricks.ErrUnimplemented
}
