package db_connection

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func Connection(ctx context.Context) *sqlx.DB {
	dsn := "host=0.0.0.0 port=5432 user=root password=root dbname=postgres sslmode=disable"
	connect, err := sqlx.ConnectContext(ctx, "postgres", dsn)
	if err != nil {
		log.Error().Err(err).Msg("error while connection to db")
	}

	if err = connect.Ping(); err != nil {
		log.Error().Msgf("error while ping to database, err: %v", err.Error())
	}

	return connect
}
