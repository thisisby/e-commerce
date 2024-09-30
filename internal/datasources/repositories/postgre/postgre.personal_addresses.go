package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
)

type postgrePersonalAddresses struct {
	Conn *sqlx.DB
}

func NewPostgrePersonalAddresses(Conn *sqlx.DB) domains.PersonalAddressesRepository {
	return &postgrePersonalAddresses{
		Conn: Conn,
	}
}

func (p *postgrePersonalAddresses) FindAll() ([]domains.PersonalAddressesDomain, error) {
	query := `
		SELECT 
			pa.id "id", 
			pa.user_id "user_id", 
			pa.street "street", 
			pa.region "region", 
			pa.apartment "apartment", 
			pa.street_num "street_num", 
			pa.city_id "city_id",
			c.id "city.id",
			c.name "city.name"
		FROM 
		personal_addresses pa
		JOIN cities c on pa.city_id = c.id
	`

	var personalAddresses []records.PersonalAddresses

	err := p.Conn.Select(&personalAddresses, query)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfPersonalAddressesDomain(personalAddresses), nil
}

func (p *postgrePersonalAddresses) FindByUserId(userId int) ([]domains.PersonalAddressesDomain, error) {
	query := `
		SELECT 
			pa.id "id", 
			pa.user_id "user_id", 
			pa.street "street", 
			pa.region "region", 
			pa.apartment "apartment", 
			pa.street_num "street_num", 
			pa.city_id "city_id",
			c.id "city.id",
			c.name "city.name"
		FROM 
		personal_addresses pa
		JOIN cities c on pa.city_id = c.id
		WHERE pa.user_id = $1
	`

	var personalAddresses []records.PersonalAddresses

	err := p.Conn.Select(&personalAddresses, query, userId)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfPersonalAddressesDomain(personalAddresses), nil
}

func (p *postgrePersonalAddresses) Save(domain domains.PersonalAddressesDomain) error {
	query := `
		INSERT INTO personal_addresses (user_id, street, region, apartment, street_num, city_id) 
		VALUES (:user_id, :street, :region, :apartment, :street_num, :city_id)
	`

	personalAddressesRecord := records.FromPersonalAddressesDomain(domain)

	_, err := p.Conn.NamedQuery(query, personalAddressesRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgrePersonalAddresses) Update(domain domains.PersonalAddressesDomain, id int) error {
	query := `
		UPDATE personal_addresses 
		SET user_id = :user_id, street = :street, region = :region, apartment = :apartment, street_num = :street_num, city_id = :city_id 
		WHERE id = :id
	`

	personalAddressesRecord := records.FromPersonalAddressesDomain(domain)
	personalAddressesRecord.Id = id

	_, err := p.Conn.NamedQuery(query, personalAddressesRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgrePersonalAddresses) Delete(id int) error {
	query := `DELETE FROM personal_addresses WHERE id = $1`

	_, err := p.Conn.Exec(query, id)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}
