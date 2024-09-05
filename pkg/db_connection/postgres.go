package dbConnection

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mrKrabsmr/test-task-verba-group/configs"
)

func PostgreSQLConnection(config configs.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", config.DBAddress)
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	if err = db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}
