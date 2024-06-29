package utils

import (
	"fmt"
	"ga_marketplace/internal/config"
	"ga_marketplace/internal/datasources/drivers"
	"github.com/jmoiron/sqlx"
	"time"
)

func SetupPostgreConnection() (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.AppConfig.DBHost, config.AppConfig.DBPort, config.AppConfig.DBUser, config.AppConfig.DBPassword, config.AppConfig.DBName,
	)

	SQLXDriver := drivers.NewSQLXDriver(
		"postgres",
		dsn,
		10,
		5,
		15*time.Minute,
	)

	conn, err := SQLXDriver.OpenConnection()
	if err != nil {
		return nil, err
	}

	return conn, nil

}
