package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
)

type postgreSubservicesRepository struct {
	Conn *sqlx.DB
}

func NewPostgreSubservicesRepository(Conn *sqlx.DB) domains.SubServicesRepository {
	return &postgreSubservicesRepository{
		Conn: Conn,
	}
}

func (p *postgreSubservicesRepository) FindAll() ([]domains.SubServicesDomain, error) {
	query := `SELECT id, name, service_id FROM subservices`

	var subservices []records.SubServiceRecord

	err := p.Conn.Select(&subservices, query)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfSubServicesDomain(subservices), nil
}

func (p *postgreSubservicesRepository) Save(domain domains.SubServicesDomain) error {
	query := `INSERT INTO subservices (name, service_id) VALUES (:name, :service_id)`

	subservicesRecord := records.FromSubServicesDomain(&domain)

	_, err := p.Conn.NamedQuery(query, subservicesRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreSubservicesRepository) Update(domain domains.SubServicesDomain) error {
	query := `UPDATE subservices SET name = :name, service_id = :service_id WHERE id = :id`

	subservicesRecord := records.FromSubServicesDomain(&domain)

	_, err := p.Conn.NamedQuery(query, subservicesRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreSubservicesRepository) Delete(id int) error {
	query := `DELETE FROM subservices WHERE id = $1`

	_, err := p.Conn.Exec(query, id)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreSubservicesRepository) FindAllByServiceId(serviceId int) ([]domains.SubServicesDomain, error) {
	query := `SELECT id, name, service_id FROM subservices WHERE service_id = $1`

	var subservices []records.SubServiceRecord

	err := p.Conn.Select(&subservices, query, serviceId)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfSubServicesDomain(subservices), nil
}

func (p *postgreSubservicesRepository) FindById(id int) (domains.SubServicesDomain, error) {
	query := `SELECT id, name, service_id FROM subservices WHERE id = $1`

	var subservice records.SubServiceRecord

	err := p.Conn.Get(&subservice, query, id)
	if err != nil {
		return domains.SubServicesDomain{}, helpers.PostgresErrorTransform(err)
	}

	return *subservice.ToDomain(), nil
}
