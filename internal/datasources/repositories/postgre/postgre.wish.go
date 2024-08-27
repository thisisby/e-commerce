package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
	"log/slog"
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
	var query = `
		SELECT 
		    w.id, w.user_id, w.product_id, w.created_at, w.updated_at,
			p.id "product.id", p.name "product.name", p.description "product.description",
			p.ingredients "product.ingredients", p.weight "product.weight", p.c_code "product.c_code", p.ed_izm "product.ed_izm",
			p.article "product.article", p.subcategory_id "product.subcategory_id", p.brand_id "product.brand_id",
			p.price "product.price", p.image "product.image", p.images "product.images",
			p.created_at "product.created_at", p.updated_at "product.updated_at",
		    s.id "product.subcategory.id", s.name "product.subcategory.name", s.category_id "product.subcategory.category_id",
		    b.id "product.brand.id", b.name "product.brand.name",
			COALESCE(d.id, -1) "product.discount.id", COALESCE(d.product_id, 0) "product.discount.product_id", COALESCE(d.discount, 0) "product.discount.discount", COALESCE(NULLIF(d.start_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) "product.discount.start_date", COALESCE(NULLIF(d.end_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) "product.discount.end_date",
			CASE WHEN d.discount IS NOT NULL THEN p.price - (p.price * d.discount / 100) ELSE p.price END AS "product.discounted_price",
		    CASE WHEN d.discount IS NOT NULL THEN (p.price - (p.price * d.discount / 100)) * c.quantity ELSE p.price * c.quantity END AS "product.total_price",
		    CASE WHEN c.product_id IS NOT NULL THEN TRUE ELSE FALSE END AS "product.is_in_cart",
			CASE WHEN w.product_id IS NOT NULL THEN TRUE ELSE FALSE END AS "product.is_in_wishlist"
		FROM wishes w
		JOIN products p ON w.product_id = p.id
		LEFT JOIN discounts d ON p.id = d.product_id AND d.start_date <= NOW() AND d.end_date >= NOW()
		LEFT JOIN cart_items c ON p.id = c.product_id AND c.user_id = w.user_id
		JOIN subcategories s ON p.subcategory_id = s.id
		JOIN brands b ON p.brand_id = b.id
		WHERE w.user_id = $1
		`
	var wishRecord []records.Wish

	err := p.conn.Select(&wishRecord, query, id)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfWishesDomain(wishRecord), nil
}

func (p *postgreWishRepository) Save(wish *domains.WishDomain) error {
	query := `
		INSERT INTO wishes (user_id, product_id, created_at)
		VALUES (:user_id, :product_id, :created_at)
		`

	wishRecord := records.FromWishDomain(wish)

	_, err := p.conn.NamedQuery(query, wishRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreWishRepository) DeleteByIdAndUserId(id int, userId int) error {
	query := `
		DELETE FROM wishes
		WHERE id = $1 AND user_id = $2
		`

	_, err := p.conn.Exec(query, id, userId)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreWishRepository) FindById(id int) (*domains.WishDomain, error) {
	query := `
		SELECT w.id, w.user_id, w.product_id, w.created_at, w.updated_at,
		       p.id "product.id", p.weight "product.weight", p.name "product.name", p.description "product.description", 
		       p.price "product.price", p.created_at "product.created_at", p.updated_at "product.updated_at"
		FROM wishes w
		JOIN products p ON w.product_id = p.id
		WHERE w.id = $1
		`

	var wishRecord records.Wish
	err := p.conn.Get(&wishRecord, query, id)
	if err != nil {
		slog.Error("postgreWishRepository.FindById", err)
		return nil, helpers.PostgresErrorTransform(err)
	}

	return wishRecord.ToDomain(), nil
}

func (p *postgreWishRepository) FindByUserIdAndProductId(userId int, productId int) (*domains.WishDomain, error) {
	query := `
		SELECT 
		    w.id, w.user_id, w.product_id, w.created_at, w.updated_at,
		    p.id "product.id", p.name "product.name", p.description "product.description",
			p.ingredients "product.ingredients", p.c_code "product.c_code", p.ed_izm "product.ed_izm",
			p.article "product.article", p.weight "product.weight", p.subcategory_id "product.subcategory_id", p.brand_id "product.brand_id",
			p.price "product.price", p.image "product.image", p.images "product.images",
			p.created_at "product.created_at", p.updated_at "product.updated_at"
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
		SELECT
		    w.id, w.user_id, w.product_id, w.created_at, w.updated_at,
		    p.id "product.id", p.weight "product.weight", p.name "product.name", p.description "product.description",
			p.ingredients "product.ingredients", p.c_code "product.c_code", p.ed_izm "product.ed_izm",
			p.article "product.article", p.subcategory_id "product.subcategory_id", p.brand_id "product.brand_id",
			p.price "product.price", p.image "product.image", p.images "product.images",
			p.created_at "product.created_at", p.updated_at "product.updated_at"
		FROM wishes w
		JOIN products p ON w.product_id = p.id
		`

	var wishRecord []records.Wish
	err := p.conn.Select(&wishRecord, query)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfWishesDomain(wishRecord), nil
}
