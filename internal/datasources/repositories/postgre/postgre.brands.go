package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
)

type postgreBrandsRepository struct {
	Conn *sqlx.DB
}

func NewPostgreBrandsRepository(Conn *sqlx.DB) domains.BrandsRepository {
	return &postgreBrandsRepository{
		Conn: Conn,
	}
}

func (p *postgreBrandsRepository) FindAll() ([]domains.BrandsDomain, error) {
	query := `SELECT * FROM brands`

	var brands []records.Brands

	err := p.Conn.Select(&brands, query)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfBrandsDomain(brands), nil
}

func (p *postgreBrandsRepository) Save(domain domains.BrandsDomain) error {
	query := `INSERT INTO brands (name, info) VALUES (:name, :info)`
	brandsRecord := records.FromBrandsDomain(domain)

	_, err := p.Conn.NamedQuery(query, brandsRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreBrandsRepository) Update(domain domains.BrandsDomain) error {
	query := `UPDATE brands SET name = :name, info = :info WHERE id = :id`
	brandsRecord := records.FromBrandsDomain(domain)

	_, err := p.Conn.NamedQuery(query, brandsRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreBrandsRepository) Delete(id int) error {
	query := `DELETE FROM brands WHERE id = $1`

	_, err := p.Conn.Exec(query, id)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}
