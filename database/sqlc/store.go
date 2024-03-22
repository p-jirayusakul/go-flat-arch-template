package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/p-jirayusakul/go-flat-arch-template/pkg/config"
)

// Store defines all functions to execute db queries and transactions
type Store interface {
	Querier
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

// NewStore creates a new store
func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}

func InitDatabase(cfg config.Config) *pgxpool.Pool {

	// connect to database
	source := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable TimeZone=Asia/Bangkok", cfg.DATABASE_USER, cfg.DATABASE_PASSWORD, cfg.DATABASE_HOST, cfg.DATABASE_PORT, cfg.DATABASE_NAME)
	conn, err := pgxpool.New(context.Background(), source)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}
