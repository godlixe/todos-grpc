package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type PostgreClient struct {
	Conn *pgx.Conn
}

func NewPostgreClient() (*PostgreClient, error) {
	conn, err := pgx.Connect(context.Background(), "host=db user=postgres password=root dbname=todos port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	psqlClient := PostgreClient{
		Conn: conn,
	}

	return &psqlClient, nil
}

func (p *PostgreClient) GetConn(ctx context.Context) (*pgx.Conn, error) {
	return p.Conn, nil
}
