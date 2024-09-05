package app

import (
	"log/slog"

	"github.com/jmoiron/sqlx"
	"github.com/mrKrabsmr/test-task-verba-group/configs"
	dbConnection "github.com/mrKrabsmr/test-task-verba-group/pkg/db_connection"
)

type App struct {
	config configs.Config
	logger *slog.Logger
	db     *sqlx.DB
}

func New(config configs.Config, logger *slog.Logger) *App {
	db, err := dbConnection.PostgreSQLConnection(config)
	if err != nil {
		panic(err)
	}

	return &App{
		logger: logger,
		config: config,
		db:     db,
	}
}
