package postgre

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
)

type postgreFilialAddresses struct {
	Conn *sqlx.DB
}

func NewPostgreFilialAddresses(Conn *sqlx.DB) domains.FilialAddressesRepository {
	return &postgreFilialAddresses{
		Conn: Conn,
	}
}

func (p *postgreFilialAddresses) FindAll() ([]domains.FilialAddressesDomain, error) {
	query := `
		SELECT 
			pa.id "id", 
			pa.street "street", 
			pa.region "region", 
			pa.apartment "apartment", 
			pa.street_num "street_num", 
			pa.city_id "city_id",
			c.id "city.id",
			c.name "city.name"
		FROM 
		filial_addresses pa
		JOIN cities c on pa.city_id = c.id
	`

	var filialAddresses []records.FilialAddresses

	err := p.Conn.Select(&filialAddresses, query)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfFilialAddressesDomain(filialAddresses), nil
}

func (p *postgreFilialAddresses) FindByUserId(userId int) ([]domains.FilialAddressesDomain, error) {
	//TODO implement me
	return nil, errors.New("not used function")
}

func (p *postgreFilialAddresses) Save(domain domains.FilialAddressesDomain) error {
	query := `
		INSERT INTO filial_addresses (street, region, apartment, street_num, city_id) 
		VALUES (:street, :region, :apartment, :street_num, :city_id)
	`

	filialAddressesRecord := records.FromFilialAddressesDomain(domain)

	_, err := p.Conn.NamedQuery(query, filialAddressesRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreFilialAddresses) Update(domain domains.FilialAddressesDomain, id int) error {
	query := `
		UPDATE filial_addresses 
		SET street = :street, region = :region, apartment = :apartment, street_num = :street_num, city_id = :city_id 
		WHERE id = :id
	`

	filialAddressesRecord := records.FromFilialAddressesDomain(domain)
	filialAddressesRecord.Id = id

	_, err := p.Conn.NamedQuery(query, filialAddressesRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreFilialAddresses) Delete(id int) error {
	query := `DELETE FROM filial_addresses WHERE id = $1`

	_, err := p.Conn.Exec(query, id)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}
