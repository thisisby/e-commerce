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

func (p *postgreCartsRepository) FindAllByUserId(id int) ([]domains.CartItemsDomain, error) {
	query := `
		SELECT c.id, c.user_id, c.product_id, c.quantity, c.created_at, c.updated_at,
		       p.id "product.id", p.name "product.name", p.description "product.description", p.price "product.price", p.created_at "product.created_at", p.updated_at "product.updated_at",
			   COALESCE(d.id, 0) "product.discount.id", COALESCE(d.product_id, 0) "product.discount.product_id", COALESCE(d.discount, 0) "product.discount.discount", COALESCE(NULLIF(d.start_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) "product.discount.start_date", COALESCE(NULLIF(d.end_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) "product.discount.end_date",
			   CASE WHEN d.discount IS NOT NULL THEN p.price - (p.price * d.discount / 100) ELSE p.price END AS "product.discounted_price",
		       CASE WHEN d.discount IS NOT NULL THEN (p.price - (p.price * d.discount / 100)) * c.quantity ELSE p.price * c.quantity END AS "product.total_price"
		FROM cart_items c
		JOIN users u ON c.user_id = u.id
		JOIN products p ON c.product_id = p.id
		LEFT JOIN discounts d ON p.id = d.product_id AND d.start_date <= NOW() AND d.end_date >= NOW()
		WHERE c.user_id = $1
		`

	var cartRecord []records.CartItems

	err := p.conn.Select(&cartRecord, query, id)
	if err != nil {
		slog.Error("PostgreCartsRepository.FindAllByUserId: ", err)
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfCartItemsDomain(cartRecord), nil
}

func (p *postgreCartsRepository) Save(cart *domains.CartItemsDomain) error {
	query := `
		INSERT INTO cart_items (user_id, product_id, quantity, created_at)
		VALUES (:user_id, :product_id, :quantity, :created_at)
		`

	cartRecord := records.FromCartsDomain(cart)

	_, err := p.conn.NamedQuery(query, cartRecord)
	if err != nil {
		slog.Error("PostgreCartsRepository.Save: ", err)
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreCartsRepository) DeleteByIdAndUserId(id int, userId int) error {
	query := `
		DELETE FROM cart_items
		WHERE id = $1 AND user_id = $2
		`

	_, err := p.conn.Exec(query, id, userId)
	if err != nil {
		slog.Error("PostgreCartsRepository.DeleteByIdAndUserId: ", err)
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
		slog.Error("PostgreCartsRepository.FindById: ", err)
		return nil, helpers.PostgresErrorTransform(err)
	}

	return cartRecord.ToDomain(), nil
}

func (p *postgreCartsRepository) FindByUserIdAndProductId(userId int, productId int) (*domains.CartItemsDomain, error) {
	query := `
		SELECT c.id, c.user_id, c.product_id, c.quantity, c.created_at, c.updated_at
		FROM cart_items c
		WHERE c.user_id = $1 AND c.product_id = $2
		`

	var cartRecord records.CartItems

	err := p.conn.Get(&cartRecord, query, userId, productId)
	if err != nil {
		slog.Error("PostgreCartsRepository.FindByUserIdAndProductId: ", err)
		return nil, helpers.PostgresErrorTransform(err)
	}

	return cartRecord.ToDomain(), nil
}

func (p *postgreCartsRepository) UpdateByIdAndUserId(cart *domains.CartItemsDomain) error {
	query := `
		UPDATE cart_items
		SET quantity = :quantity
		WHERE id = :id AND user_id = :user_id
		`

	cartRecord := records.FromCartsDomain(cart)

	_, err := p.conn.NamedQuery(query, cartRecord)
	if err != nil {
		slog.Error("PostgreCartsRepository.UpdateByIdAndUserId: ", err)
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}
