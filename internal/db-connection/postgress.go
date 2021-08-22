package db_connection

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

func NewPGConnection(ctx context.Context) *pgx.Conn {
	dsn := "postgres://root:root@0.0.0.0:5432/postgres?sslmode=disable"
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}
