package db

import (
	"database/sql"
	"fmt"

	"github.com/andydevstic/boilerplate-backend/config"
	"github.com/andydevstic/boilerplate-backend/shared/constants"
)

func ConnectDb(config *config.Config) (*sql.DB, error) {
	db, err := sql.Open(constants.DbDialect, config.DbStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database connection string: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to establish database connection: %w", err)
	}

	return db, nil
}
