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

func NewPostgreProductsRepository(conn *sqlx.DB) domains.ProductsRepository {
	return &postgreProductsRepository{
		conn: conn,
	}
}

func (p *postgreProductsRepository) FindById(id int) (*domains.ProductDomain, error) {
	query := `
		SELECT id, name, description, price, stock, image, images, created_at, updated_at 
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

func (p *postgreProductsRepository) Save(product *domains.ProductDomain) error {
	query := `
		INSERT INTO products (name, description, price, image, images, stock)
		VALUES (:name, :description, :price, :image, :images, :stock)		
		`

	productRecord := records.FromProductDomain(product)

	_, err := p.conn.NamedQuery(query, productRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreProductsRepository) FindAllForMe(id int) ([]domains.ProductDomain, error) {
	query := `
		SELECT p.id, p.name, p.description, p.price, p.image, p.images, stock, p.created_at, p.updated_at,
			COALESCE(d.id, -1) "discount.id", COALESCE(d.product_id, 0) "discount.product_id", COALESCE(d.discount, 0) "discount.discount", COALESCE(NULLIF(d.start_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) "discount.start_date", COALESCE(NULLIF(d.end_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) "discount.end_date",
		CASE WHEN c.product_id IS NOT NULL THEN TRUE ELSE FALSE END AS "is_in_cart",
		CASE WHEN w.product_id IS NOT NULL THEN TRUE ELSE FALSE END AS "is_in_wishlist",
		CASE WHEN d.discount IS NOT NULL THEN p.price - (p.price * d.discount / 100) ELSE p.price END AS "discounted_price"
		FROM products p
		LEFT JOIN cart_items c ON p.id = c.product_id AND c.user_id = $1
		LEFT JOIN wishes w ON p.id = w.product_id AND w.user_id = $1
		LEFT JOIN discounts d ON p.id = d.product_id AND d.start_date <= NOW() AND d.end_date >= NOW()
		`

	var productsRecord []records.Products

	err := p.conn.Select(&productsRecord, query, id)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfProductsDomain(productsRecord), nil
}

func (p *postgreProductsRepository) UpdateById(inDom domains.ProductDomain) error {
	tx, err := p.conn.Begin()
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var productRecord records.Products
	productRecord = records.FromProductDomain(&inDom)

	updateQuery := `
		UPDATE products
		SET name = $1, description = $2, price = $3, image = $4, images = $5, stock = $6
		WHERE id = $7
	`

	_, err = tx.Exec(updateQuery, productRecord.Name, productRecord.Description, productRecord.Price, productRecord.Image, productRecord.Images, productRecord.Stock, productRecord.Id)
	if err != nil {
		tx.Rollback()
		return helpers.PostgresErrorTransform(err)
	}

	// Delete cart items where quantity exceeds the new stock
	deleteQuery := `
		DELETE FROM cart_items
		WHERE product_id = $1 AND quantity > $2
	`

	_, err = tx.Exec(deleteQuery, inDom.Id, inDom.Stock)
	if err != nil {
		tx.Rollback()
		return helpers.PostgresErrorTransform(err)
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}
