package repositories

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"myservice/internal/domain"
)

type UserRepo struct {
	conn *pgx.Conn
}

func NewDatabaseConn(conn pgx.Conn) (*UserRepo, error) {
	return &UserRepo{conn: &conn}, nil
}

// TODO: гусь
func (ur *UserRepo) InitUserTable(ctx context.Context) error {

	_, err := ur.conn.Exec(ctx,
		"CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name VARCHAR(100), created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)")
	if err != nil {
		return fmt.Errorf("cannot init table: %w", err)

	}

	return nil
}

func (ur *UserRepo) InsertUser(ctx context.Context, name string) error {
	_, err := ur.conn.Exec(
		ctx,
		"INSERT INTO users (name) VALUES ($1)",
		name,
	)
	if err != nil {
		return fmt.Errorf("cannot INSERT in table: %w", err)

	}
	return nil
}

func (ur *UserRepo) DeleteUser(ctx context.Context, id int) error {
	_, err := ur.conn.Exec(
		ctx,
		"DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("cannot DELETE: %w", err)
	}
	return nil
}

func (ur *UserRepo) UpdateUser(ctx context.Context, id int, name string) error {
	_, err := ur.conn.Exec(
		ctx,
		"UPDATE users SET name=$1 WHERE id=$2", name, id)
	if err != nil {
		return fmt.Errorf("cannot UPDATE: %w", err)
	}
	return nil
}

func (ur *UserRepo) GetUserById(ctx context.Context, id int) (domain.Users, error) {

	var res domain.Users

	rows, err := ur.conn.Query(
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
			&res.ID,
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
