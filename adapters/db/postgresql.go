package db

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/pouyanh/lookhub/settings"
)

func newPgxConnectionPool(dbSs settings.Database, _ settings.Mode) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		dbSs.Username,
		dbSs.Password,
		net.JoinHostPort(dbSs.Host, strconv.Itoa(dbSs.Port)),
		dbSs.DbName,
	)

	return pgxpool.New(context.Background(), dsn)
}
