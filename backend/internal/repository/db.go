package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB representa la conexión a la base de datos
type DB struct {
	pool *pgxpool.Pool
}

// NewDB crea una nueva conexión a la base de datos
func NewDB(databaseURL string) (*DB, error) {
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database URL: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	// Verificar conexión
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return &DB{pool: pool}, nil
}

// Close cierra la conexión a la base de datos
func (db *DB) Close() {
	db.pool.Close()
}

// Pool retorna el pool de conexiones
func (db *DB) Pool() *pgxpool.Pool {
	return db.pool
}

// SetTenant establece el tenant para RLS
func (db *DB) SetTenant(ctx context.Context, tenantID string) error {
	_, err := db.pool.Exec(ctx, "SET LOCAL app.current_tenant = $1", tenantID)
	return err
}
