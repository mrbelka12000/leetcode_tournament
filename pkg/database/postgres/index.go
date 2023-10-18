package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/mrbelka12000/leetcode_tournament/pkg/config"
)

func Connect(cfg config.Config) (*sql.DB, error) {
	conn, err := sql.Open("postgres", cfg.PGUrl)
	if err != nil {
		return nil, fmt.Errorf("sql open: %w", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping db: %w", err)
	}

	return conn, nil
}
