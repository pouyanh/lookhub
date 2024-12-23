package db

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"

	"gitlab.snapp.ir/pouyanh/lookhub/settings"
)

func newPgxConnectionPool(db settings.Database, mode settings.Mode) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		db.Username,
		db.Password,
		net.JoinHostPort(db.Host, strconv.Itoa(db.Port)),
		db.DbName,
	)

	return pgxpool.New(context.Background(), dsn)
}
