package repositories

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"myservice/internal/models"
)

func Initusertable(ctx context.Context, conn pgx.Conn) error {

	_, err := conn.Exec(ctx,
		"CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name VARCHAR(100), createdat TIMESTAMP DEFAULT CURRENT_TIMESTAMP)")
	if err != nil {
		return fmt.Errorf("cannot init table: %w", err)

	}

	return nil
}

func Inserteuser(ctx context.Context, conn pgx.Conn, name string) error {
	_, err := conn.Exec(
		ctx,
		"INSERT INTO users (name) VALUES ($1)",
		name,
	)
	if err != nil {
		return fmt.Errorf("cannot INSERT in table: %w", err)

	}
	return nil
}

func Deleteuser(ctx context.Context, conn pgx.Conn, id int) error {
	_, err := conn.Exec(
		ctx,
		"DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("cannot DELETE: %w", err)
	}
	return nil
}

func Updateuser(ctx context.Context, conn pgx.Conn, id int, name string) error {
	_, err := conn.Exec(
		ctx,
		"UPDATE users SET name=$1 WHERE id=$2", name, id)
	if err != nil {
		return fmt.Errorf("cannot UPDATE: %w", err)
	}
	return nil
}

func GetUser(ctx context.Context, conn pgx.Conn, id int) (models.UsersModel, error) {

	var res models.UsersModel

	rows, err := conn.Query(
		ctx,
		"SELECT id, name, createdat FROM users WHERE id = $1",
		id,
	)
	if err != nil {
		return res, fmt.Errorf("cannot SELECT: %w", err)
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&res.Id,
			&res.Name,
			&res.CreatedAt,
		)
		if err != nil {
			return res, fmt.Errorf("cannot scan row: %w", err)
		}
	}

	if err := rows.Err(); err != nil {
		return res, fmt.Errorf("rows error: %w", err)
	}

	return res, nil
}
