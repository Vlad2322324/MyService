package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

func DBconnect(ctx context.Context, dbURL string) *pgx.Conn {

	conn, err := pgx.Connect(ctx, dbURL)

	if err != nil {
		log.Fatal("Error connecting to database%v\n", err)
	}

	return conn
}
