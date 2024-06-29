package seeders

import (
	"errors"
	"ga_marketplace/internal/datasources/records"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type Seeder interface {
	RolesSeeder(rolesData []records.Roles) (err error)
}

type seeder struct {
	conn *sqlx.DB
}

func NewSeeder(conn *sqlx.DB) Seeder {
	return &seeder{conn: conn}
}

func (s *seeder) RolesSeeder(rolesData []records.Roles) (err error) {
	query := `INSERT INTO roles (id, name) VALUES (:id, :name)`
	if len(rolesData) == 0 {
		return errors.New("roles data is empty")
	}

	slog.Info("Seeding roles data...")
	for _, role := range rolesData {
		_, err = s.conn.NamedQuery(query, role)
		if err != nil {
			return err
		}
	}
	slog.Info("Roles data seeded successfully")

	return nil
}
