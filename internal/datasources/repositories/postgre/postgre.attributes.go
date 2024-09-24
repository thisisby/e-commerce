package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
)

type postgreAttributesRepository struct {
	Conn *sqlx.DB
}

func NewPostgreAttributesRepository(Conn *sqlx.DB) domains.AttributesRepository {
	return &postgreAttributesRepository{
		Conn: Conn,
	}
}

func (p *postgreAttributesRepository) FindAll() ([]domains.AttributesDomain, error) {
	query := `SELECT id, name, characteristic_id FROM attributes`

	var attributes []records.Attributes

	err := p.Conn.Select(&attributes, query)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfAttributesDomain(attributes), nil

}

func (p *postgreAttributesRepository) Save(domain domains.AttributesDomain) error {
	query := `INSERT INTO attributes (name, characteristic_id) VALUES (:name, :characteristic_id)`

	attributesRecord := records.FromAttributesDomain(domain)

	_, err := p.Conn.NamedQuery(query, attributesRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreAttributesRepository) Update(domain domains.AttributesDomain) error {
	query := `UPDATE attributes SET name = :name, characteristic_id = :characteristic_id WHERE id = :id`

	attributesRecord := records.FromAttributesDomain(domain)

	_, err := p.Conn.NamedQuery(query, attributesRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreAttributesRepository) Delete(id int) error {
	query := `DELETE FROM attributes WHERE id = $1`

	_, err := p.Conn.Exec(query, id)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreAttributesRepository) FindById(id int) (domains.AttributesDomain, error) {
	query := `SELECT id, name, characteristic_id FROM attributes WHERE id = $1`

	var attributes records.Attributes

	err := p.Conn.Get(&attributes, query, id)
	if err != nil {
		return domains.AttributesDomain{}, helpers.PostgresErrorTransform(err)
	}

	return attributes.ToDomain(), nil
}

func (p *postgreAttributesRepository) FindAllByCharacteristicsId(characteristicsId int) ([]domains.AttributesDomain, error) {
	query := `SELECT id, name, characteristic_id FROM attributes WHERE characteristic_id = $1`

	var attributes []records.Attributes

	err := p.Conn.Select(&attributes, query, characteristicsId)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfAttributesDomain(attributes), nil
}
