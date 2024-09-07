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
	tx, err := p.conn.Beginx()
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	stockQuery := `
		INSERT INTO product_stock (transaction_id, customer_id, date, active)
		VALUES (:transaction_id, :customer_id, :date, :active)
	`
	productStockRecord := records.FromProductStockDomain(productStock)

	_, err = tx.NamedExec(stockQuery, productStockRecord)
	if err != nil {
		tx.Rollback()
		return helpers.PostgresErrorTransform(err)
	}

	itemQuery := `
		INSERT INTO product_stock_item (product_code, transaction_id, quantity, amount, transaction_type)
		VALUES (:product_code, :transaction_id, :quantity, :amount, :transaction_type)
	`

	for _, item := range productStock.Items {
		itemRecord := records.FromProductStockItemDomain(item)
		_, err = tx.NamedExec(itemQuery, itemRecord)
		if err != nil {
			tx.Rollback()
			return helpers.PostgresErrorTransform(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreProductStockRepository) Update(productStock domains.ProductStockDomain, transactionId string) error {
	query := `
		UPDATE product_stock 
		SET 
		    transaction_id = $1, 
		    customer_id = $2,
		    date = $3,
		    active = $4
		WHERE transaction_id = $5
	`

	productStockRecord := records.FromProductStockDomain(productStock)

	_, err := p.conn.Exec(
		query,
		productStockRecord.TransactionId,
		productStockRecord.CustomerId,
		productStockRecord.Date,
		productStockRecord.Active,
		transactionId,
	)
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

func (p *postgreProductStockRepository) FindStockItem(transactionId string, productId string) (domains.ProductStockItemDomain, error) {
	query := `
		SELECT * FROM product_stock_item WHERE transaction_id = $1 AND product_code = $2
	`

	var productStockItemRecord records.ProductStockItem
	err := p.conn.Get(&productStockItemRecord, query, transactionId, productId)
	if err != nil {
		return domains.ProductStockItemDomain{}, helpers.PostgresErrorTransform(err)
	}
	return productStockItemRecord.ToDomain(), nil
}

func (p *postgreProductStockRepository) UpdateProductStockItem(item domains.ProductStockItemDomain, transactionId string, productId string) error {
	query := `
		UPDATE product_stock_item 
		SET 
		    product_code = $1, 
		    transaction_id = $2,
		    quantity = $3,
		    amount = $4,
		    transaction_type = $5
		WHERE transaction_id = $6 AND product_code = $7
	`

	itemRecord := records.FromProductStockItemDomain(item)

	_, err := p.conn.Exec(
		query,
		itemRecord.ProductCode,
		itemRecord.TransactionId,
		itemRecord.Quantity,
		itemRecord.Amount,
		itemRecord.TransactionType,
		transactionId,
		productId,
	)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}
