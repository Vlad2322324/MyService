package repositories

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

// NewPostgresPool устанавливает пул соединений с PostgreSQL.
func NewPostgresPool(connString string) (*pgxpool.Pool, error) {
	// Парсим строку подключения для дальнейшей конфигурации
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("не удалось разобрать строку подключения: %w", err)
	}

	// Настройка пула соединений (Connection Pool)
	// Это позволяет переиспользовать соединения, повышая производительность[reference:3]
	config.MaxConns = 25                      // Максимум открытых соединений[reference:4]
	config.MinConns = 5                       // Минимум всегда открытых соединений[reference:5]
	config.MaxConnLifetime = time.Hour        // Время жизни соединения[reference:6]
	config.MaxConnIdleTime = 30 * time.Minute // Время жизни простаивающего соединения

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать пул соединений: %w", err)
	}

	// Проверяем подключение, пингуя базу данных
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := pool.Ping(ctx); err != nil {
		// Если пинг не удался, закрываем пул перед возвратом ошибки
		pool.Close()
		return nil, fmt.Errorf("база данных недоступна: %w", err)
	}

	log.Println("Успешно подключено к PostgreSQL")
	return pool, nil
}
