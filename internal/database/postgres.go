package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func DBconnect(ctx context.Context, dbURL string) (*pgx.Conn, error) {

	conn, err := pgx.Connect(ctx, dbURL)

	if err != nil {
		return nil, fmt.Errorf("Error connecting to database%v\n", err)
	}

	return conn, nil
}
