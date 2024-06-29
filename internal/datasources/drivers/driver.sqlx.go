package drivers

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

type SQLXDriver struct {
	DriverName      string
	DataSourceName  string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func NewSQLXDriver(driverName string, dataSourceName string, maxOpenConns int, maxIdleConns int, connMaxLifetime time.Duration) *SQLXDriver {
	return &SQLXDriver{
		DriverName:      driverName,
		DataSourceName:  dataSourceName,
		MaxOpenConns:    maxOpenConns,
		MaxIdleConns:    maxIdleConns,
		ConnMaxLifetime: connMaxLifetime,
	}
}

func (d *SQLXDriver) OpenConnection() (*sqlx.DB, error) {
	db, err := sqlx.Open(d.DriverName, d.DataSourceName)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(d.MaxOpenConns)
	db.SetMaxIdleConns(d.MaxIdleConns)
	db.SetConnMaxLifetime(d.ConnMaxLifetime)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
