package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
)

type postgreServicesRepository struct {
	Conn *sqlx.DB
}

func NewPostgreServicesRepository(Conn *sqlx.DB) domains.ServicesRepository {
	return &postgreServicesRepository{
		Conn: Conn,
	}
}

func (p *postgreServicesRepository) FindAll() ([]domains.ServicesDomain, error) {
	query := `SELECT id, name FROM services`

	var services []records.Services

	err := p.Conn.Select(&services, query)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfServicesDomain(services), nil
}

func (p *postgreServicesRepository) Save(domain domains.ServicesDomain) error {
	query := `INSERT INTO services (name) VALUES (:name)`
	servicesRecord := records.FromServicesDomain(domain)

	_, err := p.Conn.NamedQuery(query, servicesRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreServicesRepository) Update(domain domains.ServicesDomain) error {
	query := `UPDATE services SET name = :name WHERE id = :id`
	servicesRecord := records.FromServicesDomain(domain)

	_, err := p.Conn.NamedQuery(query, servicesRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreServicesRepository) Delete(id int) error {
	query := `DELETE FROM services WHERE id = $1`

	_, err := p.Conn.Exec(query, id)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}
