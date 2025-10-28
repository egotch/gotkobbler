package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/egotch/gotkobbler/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func New(cfg *config.Config) (*DB, error) {
	ctx := context.Background()

	// configure the connection pool
	poolConfig, err := pgxpool.ParseConfig(cfg.DatabaseURL())
	if err != nil {
		return nil, fmt.Errorf("unable to parse database config: %w", err)
	}

	poolConfig.MaxConns = 25
	poolConfig.MinConns = 5
	poolConfig.MaxConnLifetime = time.Hour
	poolConfig.MaxConnIdleTime = 30 * time.Minute
	poolConfig.HealthCheckPeriod = 1 * time.Minute

	// connect to database
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	// verify successful connection
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("unable to ping the database: %w", err)
	}	

	log.Println("✅ Successfully connected to database")

	return &DB{Pool: pool}, nil
}

// close the database connection and exit gracefully
func (db *DB) Close() {
	db.Pool.Close()

	log.Println("📪 Database connection is now closed.")
}

func (db *DB) Ping(ctx context.Context) error {
	return db.Pool.Ping(ctx)
}
