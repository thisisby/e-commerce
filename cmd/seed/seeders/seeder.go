package seeders

import (
	"errors"
	"ga_marketplace/internal/datasources/records"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type Seeder interface {
	RolesSeeder(rolesData []records.Roles) (err error)
	ProductsSeeder(productsData []records.Products) (err error)
	UsersSeeder(usersData []records.Users) (err error)
	CitiesSeeder(citiesData []records.Cities) (err error)
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

func (s *seeder) ProductsSeeder(productsData []records.Products) (err error) {
	query := `
        INSERT INTO products (name, description, price, image, stock, created_at, updated_at)
        VALUES (:name, :description, :price, :image, :stock, :created_at, :updated_at)
    `
	if len(productsData) == 0 {
		return errors.New("products data is empty")
	}

	slog.Info("Seeding products data...")
	for _, product := range productsData {
		_, err = s.conn.NamedExec(query, product)
		if err != nil {
			return err
		}
	}
	slog.Info("Products data seeded successfully")

	return nil
}

func (s *seeder) UsersSeeder(usersData []records.Users) (err error) {
	query := `
		INSERT INTO users (name, phone, role_id, refresh_token, date_of_birth, created_at, updated_at)
		VALUES (:name, :phone, :role_id, :refresh_token, :date_of_birth, :created_at, :updated_at)
	`
	if len(usersData) == 0 {
		return errors.New("users data is empty")
	}

	slog.Info("Seeding users data...")
	for _, user := range usersData {
		_, err = s.conn.NamedExec(query, user)
		if err != nil {
			return err
		}
	}
	slog.Info("Users data seeded successfully")

	return nil
}

func (s *seeder) CitiesSeeder(citiesData []records.Cities) (err error) {
	query := `INSERT INTO cities (id, name) VALUES (:id, :name)`
	if len(citiesData) == 0 {
		return errors.New("city data is empty")
	}

	slog.Info("Seeding Cities data...")
	for _, country := range citiesData {
		_, err = s.conn.NamedQuery(query, country)
		if err != nil {
			return err
		}
	}
	slog.Info("Cities data seeded successfully")

	return nil
}
