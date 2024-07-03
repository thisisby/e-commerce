package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
)

type postgreWishRepository struct {
	conn *sqlx.DB
}

func NewPostgreWishRepository(conn *sqlx.DB) domains.WishRepository {
	return &postgreWishRepository{
		conn: conn,
	}
}

func (p *postgreWishRepository) FindByUserId(id int) ([]domains.WishDomain, error) {
	query := `
		SELECT w.id, w.user_id, w.product_id, w.created_at, w.updated_at,
		       p.id "product.id", p.name "product.name", p.description "product.description", p.price "product.price", p.created_at "product.created_at", p.updated_at "product.updated_at"
		FROM wishes w
		JOIN products p ON w.product_id = p.id
		WHERE w.user_id = $1
		`

	var wishRecord []records.Wish
	err := p.conn.Select(&wishRecord, query, id)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfWishDomain(wishRecord), nil
}

func (p *postgreWishRepository) Save(wish *domains.WishDomain) error {
	query := `
		INSERT INTO wishes (user_id, product_id, created_at)
		VALUES (:user_id, :product_id, :created_at)`

	wishRecord := records.FromWishDomain(wish)

	_, err := p.conn.NamedQuery(query, wishRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreWishRepository) Delete(id int, userId int) error {
	query := `
		DELETE FROM wishes
		WHERE id = $1 AND user_id = $2`

	_, err := p.conn.Exec(query, id, userId)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreWishRepository) FindById(id int) (*domains.WishDomain, error) {
	query := `
		SELECT w.id, w.user_id, w.product_id, w.created_at, w.updated_at,
		       p.id "product.id", p.name "product.name", p.description "product.description", p.price "product.price", p.created_at "product.created_at", p.updated_at "product.updated_at"
		FROM wishes w
		JOIN products p ON w.product_id = p.id
		WHERE w.id = $1
		`

	var wishRecord records.Wish
	err := p.conn.Get(&wishRecord, query, id)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return wishRecord.ToDomain(), nil
}

func (p *postgreWishRepository) FindByUserIdAndProductId(userId int, productId int) (*domains.WishDomain, error) {
	query := `
		SELECT w.id, w.user_id, w.product_id, w.created_at, w.updated_at,
		       p.id "product.id", p.name "product.name", p.description "product.description", p.price "product.price", p.created_at "product.created_at", p.updated_at "product.updated_at"
		FROM wishes w
		JOIN products p ON w.product_id = p.id
		WHERE w.user_id = $1 AND w.product_id = $2
		`

	var wishRecord records.Wish
	err := p.conn.Get(&wishRecord, query, userId, productId)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return wishRecord.ToDomain(), nil
}

func (p *postgreWishRepository) FindAll() ([]domains.WishDomain, error) {
	query := `
		SELECT w.id, w.user_id, w.product_id, w.created_at, w.updated_at,
		       p.id "product.id", p.name "product.name", p.description "product.description", p.price "product.price", p.created_at "product.created_at", p.updated_at "product.updated_at"
		FROM wishes w
		JOIN products p ON w.product_id = p.id
		`

	var wishRecord []records.Wish
	err := p.conn.Select(&wishRecord, query)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfWishDomain(wishRecord), nil
}
