package sql

import (
	"context"
	"fmt"
	"test/v2/internal/types"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pool *pgxpool.Pool
	err  error
)

func PgPoolInit() (*pgxpool.Pool, error) {
	psqlconn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=require&pool_max_conns=%d",
		types.DB_user, types.DB_password, types.DB_host, types.DB_port, types.DB_name, types.DB_max_conn)

	ctx, cancel := context.WithTimeout(context.Background(), types.DB_timeout)
	defer cancel()
	pool, err = pgxpool.New(ctx, psqlconn)
	if err != nil {
		return nil, fmt.Errorf("postgreSQL init pool error: %v", err)
	}
	return pool, nil
}

func PgPoolClose() {
	pool.Close()
}
