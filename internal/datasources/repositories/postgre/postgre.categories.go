package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
)

type postgreCategoriesRepository struct {
	Conn *sqlx.DB
}

func NewPostgreCategoriesRepository(Conn *sqlx.DB) domains.CategoriesRepository {
	return &postgreCategoriesRepository{
		Conn: Conn,
	}
}

func (p *postgreCategoriesRepository) FindAll() ([]domains.CategoriesDomain, error) {
	query := `SELECT id, name FROM categories`

	var categories []records.Categories

	err := p.Conn.Select(&categories, query)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfCategoriesDomain(categories), nil
}

func (p *postgreCategoriesRepository) Save(domain domains.CategoriesDomain) error {
	query := `INSERT INTO categories (name) VALUES (:name)`
	categoriesRecord := records.FromCategoriesDomain(domain)

	_, err := p.Conn.NamedQuery(query, categoriesRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreCategoriesRepository) Update(domain domains.CategoriesDomain) error {
	query := `UPDATE categories SET name = :name WHERE id = :id`
	categoriesRecord := records.FromCategoriesDomain(domain)

	_, err := p.Conn.NamedQuery(query, categoriesRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreCategoriesRepository) Delete(id int) error {
	query := `DELETE FROM categories WHERE id = $1`

	_, err := p.Conn.Exec(query, id)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}
