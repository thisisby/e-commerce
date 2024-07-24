package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type postgreSubcategoriesRepository struct {
	Conn *sqlx.DB
}

func NewPostgreSubcategoriesRepository(Conn *sqlx.DB) domains.SubcategoriesRepository {
	return &postgreSubcategoriesRepository{
		Conn: Conn,
	}
}

func (p *postgreSubcategoriesRepository) FindAll() ([]domains.SubcategoriesDomain, error) {
	query := `SELECT id, name, category_id FROM subcategories`

	var subcategories []records.SubcategoriesRecord

	err := p.Conn.Select(&subcategories, query)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfSubcategoriesDomain(subcategories), nil
}

func (p *postgreSubcategoriesRepository) FindAllByCategoryId(categoryId int) ([]domains.SubcategoriesDomain, error) {
	query := `SELECT id, name, category_id FROM subcategories WHERE category_id = $1`

	var subcategories []records.SubcategoriesRecord

	err := p.Conn.Select(&subcategories, query, categoryId)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfSubcategoriesDomain(subcategories), nil
}

func (p *postgreSubcategoriesRepository) Save(subcategoriesDomain domains.SubcategoriesDomain) error {
	query := `INSERT INTO subcategories (name, category_id) VALUES (:name, :category_id)`

	subcategoriesRecord := records.FromSubcategoriesDomain(&subcategoriesDomain)

	_, err := p.Conn.NamedQuery(query, subcategoriesRecord)
	if err != nil {
		slog.Error("postgreSubcategoriesRepository.Save - p.Conn.NamedQuery: ", err)
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreSubcategoriesRepository) Update(subcategoriesDomain domains.SubcategoriesDomain) error {
	query := `UPDATE subcategories SET name = :name, category_id = :category_id WHERE id = :id`

	subcategoriesRecord := records.FromSubcategoriesDomain(&subcategoriesDomain)

	_, err := p.Conn.NamedQuery(query, subcategoriesRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreSubcategoriesRepository) Delete(id int) error {
	query := `DELETE FROM subcategories WHERE id = $1`

	_, err := p.Conn.Exec(query, id)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreSubcategoriesRepository) FindById(id int) (domains.SubcategoriesDomain, error) {
	query := `SELECT id, name, category_id FROM subcategories WHERE id = $1`

	var subcategories records.SubcategoriesRecord

	err := p.Conn.Get(&subcategories, query, id)
	if err != nil {
		return domains.SubcategoriesDomain{}, helpers.PostgresErrorTransform(err)
	}

	return *subcategories.ToDomain(), nil
}
