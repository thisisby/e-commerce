package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
)

type postgreDiscountRepository struct {
	conn *sqlx.DB
}

func NewPostgreDiscountRepository(conn *sqlx.DB) domains.DiscountsRepository {
	return &postgreDiscountRepository{
		conn: conn,
	}
}

func (p *postgreDiscountRepository) Save(discount *domains.DiscountsDomain) error {
	query := `
		INSERT INTO discounts (product_id, discount, start_date, end_date) 
		VALUES (:product_id, :discount, :start_date, :end_date)
		`

	discountRecord := records.FromDiscountsDomain(discount)

	_, err := p.conn.NamedQuery(query, discountRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreDiscountRepository) DeleteByProductId(id int) error {
	query := `
		DELETE FROM discounts
		WHERE product_id = $1
		`

	_, err := p.conn.Exec(query, id)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreDiscountRepository) FindByProductId(productId int) (*domains.DiscountsDomain, error) {
	query := `
		SELECT * FROM discounts
		WHERE product_id = $1
		`

	var discountRecord records.Discounts

	err := p.conn.Get(&discountRecord, query, productId)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return discountRecord.ToDiscountsDomain(), nil
}
