package database

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type PostgreItf interface {

	// GetConn returns a database connection
	GetConn(ctx context.Context) (*pgx.Conn, error)
}
