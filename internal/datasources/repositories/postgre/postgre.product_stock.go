package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
)

type postgreProductStockRepository struct {
	conn *sqlx.DB
}

func NewPostgreProductStockRepository(conn *sqlx.DB) domains.ProductStockRepository {
	return &postgreProductStockRepository{conn}
}

func (p *postgreProductStockRepository) Save(productStock domains.ProductStockDomain) error {
	query := `
		INSERT INTO product_stock (c_code, date, transaction_type, transaction_id, quantity, total_sum, transaction_status)
		VALUES(:c_code, :date, :transaction_type, :transaction_id, :quantity, :total_sum, :transaction_status)		
	`

	productStockRecord := records.FromProductStockDomain(productStock)

	_, err := p.conn.NamedExec(query, productStockRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreProductStockRepository) Update(productStock domains.ProductStockDomain) error {
	query := `
		UPDATE product_stock 
		SET 
		    c_code = :c_code, 
		    date = :date,
		    transaction_type = :transaction_type,
		    transaction_id = :transaction_id,
		    quantity = :quantity,
		    total_sum = :total_sum,
		    transaction_status = :transaction_status
		WHERE id = :id
	`

	productStockRecord := records.FromProductStockDomain(productStock)

	_, err := p.conn.NamedExec(query, productStockRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreProductStockRepository) FindById(id string) (domains.ProductStockDomain, error) {
	query := `
		SELECT * FROM product_stock WHERE transaction_id = $1
	`

	var productStockRecord records.ProductStock
	err := p.conn.Get(&productStockRecord, query, id)
	if err != nil {
		return domains.ProductStockDomain{}, helpers.PostgresErrorTransform(err)
	}
	return productStockRecord.ToDomain(), nil
}
