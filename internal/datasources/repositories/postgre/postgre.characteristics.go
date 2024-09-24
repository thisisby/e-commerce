package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
)

type postgreCharacteristicsRepository struct {
	Conn *sqlx.DB
}

func NewPostgreCharacteristicsRepository(Conn *sqlx.DB) domains.CharacteristicsRepository {
	return &postgreCharacteristicsRepository{
		Conn: Conn,
	}
}

func (p *postgreCharacteristicsRepository) FindAll() ([]domains.CharacteristicsDomain, error) {
	query := `SELECT id, name, subcategory_id FROM characteristics`

	var characteristics []records.Characteristics

	err := p.Conn.Select(&characteristics, query)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfCharacteristicsDomain(characteristics), nil
}

func (p *postgreCharacteristicsRepository) Save(domain domains.CharacteristicsDomain) error {
	query := `INSERT INTO characteristics (name, subcategory_id) VALUES (:name, :subcategory_id)`

	characteristicsRecord := records.FromCharacteristicsDomain(domain)

	_, err := p.Conn.NamedQuery(query, characteristicsRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreCharacteristicsRepository) Update(domain domains.CharacteristicsDomain) error {
	query := `UPDATE characteristics SET name = :name, subcategory_id = :subcategory_id WHERE id = :id`

	characteristicsRecord := records.FromCharacteristicsDomain(domain)

	_, err := p.Conn.NamedQuery(query, characteristicsRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreCharacteristicsRepository) Delete(id int) error {
	query := `DELETE FROM characteristics WHERE id = $1`

	_, err := p.Conn.Exec(query, id)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreCharacteristicsRepository) FindAllBySubcategoryId(subcategoryId int) ([]domains.CharacteristicsDomain, error) {
	query := `SELECT id, name, subcategory_id FROM characteristics WHERE subcategory_id = $1`

	var characteristics []records.Characteristics

	err := p.Conn.Select(&characteristics, query, subcategoryId)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfCharacteristicsDomain(characteristics), nil
}

func (p *postgreCharacteristicsRepository) FindById(id int) (domains.CharacteristicsDomain, error) {
	query := `SELECT id, name, subcategory_id FROM characteristics WHERE id = $1`

	var characteristics records.Characteristics

	err := p.Conn.Get(&characteristics, query, id)
	if err != nil {
		return domains.CharacteristicsDomain{}, helpers.PostgresErrorTransform(err)
	}

	return characteristics.ToDomain(), nil
}
