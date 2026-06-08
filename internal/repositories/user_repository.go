package repositories

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"myservice/internal/domain"
	"myservice/pkg/logger"
)

type UserRepo struct {
	conn *pgx.Conn
}

func NewDatabaseConn(conn *pgx.Conn) (*UserRepo, error) {
	return &UserRepo{conn: conn}, nil
}

// InitUserTable TODO: гусь
func (ur *UserRepo) InitUserTable(ctx context.Context) error {
	logger.Log.Infof("initializing users table if not exists")
	_, err := ur.conn.Exec(ctx,
		"CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name VARCHAR(100), created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)")
	if err != nil {
		return fmt.Errorf("cannot init table: %w", err)

	}

	return nil
}

func (ur *UserRepo) InsertUser(ctx context.Context, name string) (*domain.Users, error) {
	var user domain.Users
	err := ur.conn.QueryRow(
		ctx,
		"INSERT INTO users (name) VALUES ($1) RETURNING id, name, created_at",
		name,
	).Scan(&user.ID, &user.Name, &user.CreatedAt)
	if err != nil {
		logger.Log.Errorf("InsertUser failed: %v", err)
		return nil, fmt.Errorf("cannot INSERT in table: %w", err)
	}
	logger.Log.Debugf("InsertUser succeeded: id=%d name=%s", user.ID, user.Name)
	return &user, nil
}

func (ur *UserRepo) DeleteUser(ctx context.Context, id int) (int64, error) {
	tag, err := ur.conn.Exec(
		ctx,
		"DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return 0, fmt.Errorf("cannot DELETE: %w", err)
	}
	return tag.RowsAffected(), nil
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

	err := ur.conn.QueryRow(
		ctx,
		"SELECT id, name, created_at FROM users WHERE id = $1",
		id,
	).Scan(
		&res.ID,
		&res.Name,
		&res.CreatedAt,
	)
	if err != nil {
		logger.Log.Errorf("GetUserById failed (id=%d): %v", id, err)
		return res, err
	}
	logger.Log.Debugf("GetUserById succeeded: id=%d name=%s", res.ID, res.Name)

	return res, nil
}
