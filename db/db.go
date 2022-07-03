package db

import (
	"database/sql"
	"fmt"

	"github.com/andydevstic/boilerplate-backend/config"
	"github.com/andydevstic/boilerplate-backend/shared/constants"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb(config *config.Config) (*gorm.DB, error) {
	db, err := sql.Open(constants.DbDialect, config.DbStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database connection string: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to establish database connection: %w", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}))

	if err != nil {
		return nil, fmt.Errorf("connect gorm: %w", err)
	}

	return gormDB, nil
}
