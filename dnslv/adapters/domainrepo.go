package adapters

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/janstoon/toolbox/bricks"
	"github.com/janstoon/toolbox/tricks"

	"gitlab.snapp.ir/pouyanh/lookhub/dnslv"
	"gitlab.snapp.ir/pouyanh/lookhub/dnslv/adapters/db"
)

type domainRepo struct {
	conn    *pgxpool.Pool
	queries *db.Queries
}

func newDomainRepo(conn *pgxpool.Pool) domainRepo {
	return domainRepo{
		conn:    conn,
		queries: db.New(conn),
	}
}

func (repo domainRepo) GetUnexpiredDomain(ctx context.Context, name string, ttl time.Duration) (*dnslv.Domain, error) {
	dst, err := dnslv.DomainByName(dnslv.DomainName(name))
	if err != nil {
		return nil, err
	}

	rows, err := repo.queries.GetUnexpiredDomain(ctx, db.GetUnexpiredDomainParams{
		Name: name,
		Ttl:  int32(ttl.Seconds()),
	})
	if err != nil {
		return nil, err
	}

	err = dst.AddRecords(tricks.Map(rows, func(src db.GetUnexpiredDomainRow) dnslv.ResourceRecord {
		return dnslv.ResourceRecord{
			Type:  dnslv.ResourceRecordType(tricks.PtrVal(src.ResourceRecordType)),
			Value: tricks.PtrVal(src.ResourceRecordValue),
			TTL:   dnslv.ResourceRecordTTL(tricks.PtrVal(src.ResourceRecordTtl)),
		}
	})...)
	if err != nil {
		return nil, err
	}

	return dst, nil
}

func (repo domainRepo) SyncDomain(ctx context.Context, domain *dnslv.Domain) error {
	tx, err := repo.conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback(context.Background()) }()

	qtx := repo.queries.WithTx(tx)

	daoDomain, err := qtx.UpsertDomain(ctx, domain.Name().String())
	if err != nil {
		return err
	}

	err = qtx.DeleteDomainResourceRecordsByDomainID(ctx, tricks.ValPtr(daoDomain.ID))
	if err != nil {
		return err
	}

	count, err := qtx.AddResourceRecord(ctx,
		tricks.Map(domain.Records(), func(src dnslv.ResourceRecord) db.AddResourceRecordParams {
			return db.AddResourceRecordParams{
				DomainID: tricks.ValPtr(daoDomain.ID),
				Type:     string(src.Type),
				Value:    src.Value,
				Ttl:      int32(src.TTL),
			}
		}),
	)
	if err != nil {
		return err
	}

	if count != int64(len(domain.Records())) {
		return bricks.ErrDataLoss
	}

	return tx.Commit(ctx)
}
