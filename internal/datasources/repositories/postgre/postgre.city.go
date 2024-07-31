package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
)

type postgreCityRepository struct {
	Conn *sqlx.DB
}

func NewPostgreCityRepository(Conn *sqlx.DB) domains.CitiesRepository {
	return &postgreCityRepository{
		Conn: Conn,
	}
}

func (p *postgreCityRepository) FindAll() ([]domains.CityDomain, error) {
	query := `
		SELECT id, name, delivery_duration_days
		FROM cities
	`

	var cities []records.Cities

	err := p.Conn.Select(&cities, query)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfCityDomain(cities), nil
}

func (p *postgreCityRepository) FindById(id int) (domains.CityDomain, error) {
	query := `
		SELECT id, name, delivery_duration_days
		FROM cities
		WHERE id = $1
	`

	var city records.Cities

	err := p.Conn.Get(&city, query, id)
	if err != nil {
		return domains.CityDomain{}, helpers.PostgresErrorTransform(err)
	}

	return *city.ToDomain(), nil
}

func (p *postgreCityRepository) Save(city domains.CityDomain) error {
	query := `INSERT INTO cities (name, delivery_duration_days) VALUES (:name, :delivery_duration_days)`
	cityRecord := records.FromCityDomain(&city)

	_, err := p.Conn.NamedQuery(query, cityRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreCityRepository) Update(city domains.CityDomain) error {
	query := `UPDATE cities SET name = :name, delivery_duration_days = :delivery_duration_days WHERE id = :id`
	cityRecord := records.FromCityDomain(&city)

	_, err := p.Conn.NamedQuery(query, cityRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreCityRepository) Delete(id int) error {
	query := `DELETE FROM cities WHERE id = $1`

	_, err := p.Conn.Exec(query, id)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}
