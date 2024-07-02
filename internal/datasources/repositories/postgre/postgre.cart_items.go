package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type postgreCartsRepository struct {
	conn *sqlx.DB
}

func NewPostgreCartsRepository(conn *sqlx.DB) domains.CartItemsRepository {
	return &postgreCartsRepository{
		conn: conn,
	}
}

func (p *postgreCartsRepository) FindByUserId(id int) ([]domains.CartItemsDomain, error) {
	query := `
		SELECT c.id, c.user_id, c.product_id, c.quantity, c.created_at, c.updated_at,
		       p.id "product.id", p.name "product.name", p.description "product.description", p.price "product.price", p.created_at "product.created_at", p.updated_at "product.updated_at"
		FROM cart_items c
		JOIN users u ON c.user_id = u.id
		JOIN products p ON c.product_id = p.id
		WHERE c.user_id = $1
		`

	var cartRecord []records.CartItems

	err := p.conn.Select(&cartRecord, query, id)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfDomain(cartRecord), nil
}

func (p *postgreCartsRepository) Save(cart *domains.CartItemsDomain) error {
	query := `
		INSERT INTO cart_items (user_id, product_id, quantity, created_at)
		VALUES (:user_id, :product_id, :quantity, :created_at)`

	cartRecord := records.FromCartsDomain(cart)

	_, err := p.conn.NamedQuery(query, cartRecord)
	if err != nil {
		slog.Error("PostgreCartsRepository.Save: ", err)
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreCartsRepository) Delete(id int, userId int) error {
	query := `
		DELETE FROM cart_items
		WHERE id = $1 AND user_id = $2
		`

	_, err := p.conn.Exec(query, id, userId)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreCartsRepository) FindById(id int) (*domains.CartItemsDomain, error) {
	query := `
		SELECT c.id, c.user_id, c.product_id, c.quantity, c.created_at, c.updated_at,
		       p.id "product.id", p.name "product.name", p.description "product.description", p.price "product.price", p.created_at "product.created_at", p.updated_at "product.updated_at"
		FROM cart_items c
		JOIN users u ON c.user_id = u.id
		JOIN products p ON c.product_id = p.id
		WHERE c.id = $1
		`

	var cartRecord records.CartItems

	err := p.conn.Get(&cartRecord, query, id)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return cartRecord.ToDomain(), nil
}

func (p *postgreCartsRepository) FindAll() ([]domains.CartItemsDomain, error) {
	query := `
		SELECT c.id, c.user_id, c.product_id, c.quantity, c.created_at, c.updated_at,
		       p.id "product.id", p.name "product.name", p.description "product.description", p.price "product.price", p.created_at "product.created_at", p.updated_at "product.updated_at"
		FROM cart_items c
		JOIN users u ON c.user_id = u.id
		JOIN products p ON c.product_id = p.id
		`

	var cartRecord []records.CartItems

	err := p.conn.Select(&cartRecord, query)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfDomain(cartRecord), nil
}
