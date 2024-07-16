package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
)

type postgreCountriesRepository struct {
	conn *sqlx.DB
}

func NewPostgreCountriesRepository(conn *sqlx.DB) domains.CountriesRepository {
	return &postgreCountriesRepository{
		conn: conn,
	}
}

func (p *postgreCountriesRepository) FindAll() ([]domains.CountryDomain, error) {
	query := `
		SELECT id, name
		FROM countries
		`

	var countries []records.Countries

	err := p.conn.Select(&countries, query)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfCountryDomain(countries), nil
}

func (p *postgreCountriesRepository) FindById(id int) (domains.CountryDomain, error) {
	query := `
		SELECT id, name
		FROM countries
		WHERE id = $1
		`

	var country records.Countries

	err := p.conn.Get(&country, query, id)
	if err != nil {
		return domains.CountryDomain{}, helpers.PostgresErrorTransform(err)
	}

	return *country.ToDomain(), nil
}

func (p *postgreCountriesRepository) Save(country domains.CountryDomain) error {
	query := `INSERT INTO countries (name) VALUES (:name)`
	countryRecord := records.FromCountryDomain(&country)

	_, err := p.conn.NamedQuery(query, countryRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreCountriesRepository) Update(country domains.CountryDomain) error {
	query := `UPDATE countries SET name = :name WHERE id = :id`

	countryRecord := records.FromCountryDomain(&country)

	_, err := p.conn.NamedQuery(query, countryRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreCountriesRepository) Delete(id int) error {
	query := `DELETE FROM countries WHERE id = $1`

	_, err := p.conn.Exec(query, id)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}
