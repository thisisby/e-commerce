package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
)

type postgreProductsRepository struct {
	conn *sqlx.DB
}

func NewPostgreProductsRepository(conn *sqlx.DB) domains.ProductRepository {
	return &postgreProductsRepository{
		conn: conn,
	}
}

func (p *postgreProductsRepository) FindById(id int) (*domains.ProductDomain, error) {
	query := `
		SELECT id, name, description, price, created_at, updated_at 
		FROM products 
		WHERE id = $1
		`

	var productRecord records.Products

	err := p.conn.Get(&productRecord, query, id)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return productRecord.ToDomain(), nil
}
