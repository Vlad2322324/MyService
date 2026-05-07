package repositories

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func Initusertable(ctx context.Context, conn pgx.Conn) error {

	_, err := conn.Exec(ctx,
		"CREATE TABLE IF NOT EXISTS stub (id SERIAL PRIMARY KEY)")
	if err != nil {
		return fmt.Errorf("cannot init table: %w", err)

	}

	return nil
}
