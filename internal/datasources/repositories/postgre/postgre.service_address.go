package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
)

type postgreServiceAddressRepository struct {
	conn *sqlx.DB
}

func NewPostgreServiceAddressRepository(conn *sqlx.DB) domains.ServiceAddressRepository {
	return &postgreServiceAddressRepository{
		conn: conn,
	}
}

func (p *postgreServiceAddressRepository) FindAll() ([]domains.ServiceAddressDomain, error) {
	query := `SELECT s.id, s.city_id, s.address, c.id "city.id", c.name "city.name" FROM service_addresses s INNER JOIN cities c ON s.city_id = c.id`

	var serviceAddresses []records.ServiceAddress
	err := p.conn.Select(&serviceAddresses, query)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfServiceAddressDomain(serviceAddresses), nil
}

func (p *postgreServiceAddressRepository) FindById(id int) (domains.ServiceAddressDomain, error) {
	query := `SELECT s.id, s.city_id, s.address, c.id "city.id", c.name "city.name" FROM service_addresses s INNER JOIN cities c ON s.city_id = c.id WHERE s.id = $1`

	var serviceAddress records.ServiceAddress
	err := p.conn.Get(&serviceAddress, query, id)
	if err != nil {
		return domains.ServiceAddressDomain{}, helpers.PostgresErrorTransform(err)
	}

	return serviceAddress.ToDomain(), nil
}

func (p *postgreServiceAddressRepository) Save(serviceAddress domains.ServiceAddressDomain) error {
	query := `INSERT INTO service_addresses (city_id, address) VALUES (:city_id, :address)`
	serviceAddressRecord := records.FromServiceAddressDomain(serviceAddress)

	_, err := p.conn.NamedQuery(query, serviceAddressRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreServiceAddressRepository) Update(serviceAddress domains.ServiceAddressDomain) error {
	query := `UPDATE service_addresses SET city_id = :city_id, address = :address WHERE id = :id`
	serviceAddressRecord := records.FromServiceAddressDomain(serviceAddress)

	_, err := p.conn.NamedQuery(query, serviceAddressRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreServiceAddressRepository) Delete(id int) error {
	query := `DELETE FROM service_addresses WHERE id = $1`

	_, err := p.conn.Exec(query, id)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreServiceAddressRepository) FindAllByCityId(cityId int) ([]domains.ServiceAddressDomain, error) {
	query := `SELECT s.id, s.city_id, s.address, c.id "city.id", c.name "city.name" FROM service_addresses s INNER JOIN cities c ON s.city_id = c.id WHERE s.city_id = $1`

	var serviceAddresses []records.ServiceAddress
	err := p.conn.Select(&serviceAddresses, query, cityId)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfServiceAddressDomain(serviceAddresses), nil
}
