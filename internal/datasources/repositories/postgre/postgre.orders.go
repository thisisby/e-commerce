package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
)

type postgreOrdersRepository struct {
	conn *sqlx.DB
}

func NewPostgreOrdersRepository(conn *sqlx.DB) domains.OrdersRepository {
	return &postgreOrdersRepository{
		conn: conn,
	}
}

func (p *postgreOrdersRepository) Save(orders domains.OrdersDomain) error {
	tx, err := p.conn.Begin()
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	query := `
		INSERT INTO orders (user_id, total_price, discounted_price, status, street, region, apartment)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`
	var orderId int
	err = tx.QueryRow(query, orders.UserId, orders.TotalPrice, orders.DiscountedPrice, orders.Status, orders.Street, orders.Region, orders.Apartment).Scan(&orderId)
	if err != nil {
		tx.Rollback()
		return helpers.PostgresErrorTransform(err)
	}
	orders.Id = orderId

	queryDetails := `
		INSERT INTO order_details (order_id, product_id, quantity, price, sub_total)
		VALUES ($1, $2, $3, $4, $5)
	`
	for _, detail := range orders.OrderDetails {
		detail.OrderId = orderId
		_, err = tx.Exec(queryDetails, detail.OrderId, detail.ProductId, detail.Quantity, detail.Price, detail.SubTotal)
		if err != nil {
			tx.Rollback()
			return helpers.PostgresErrorTransform(err)
		}
	}

	return tx.Commit()
}

func (p *postgreOrdersRepository) FindByUserId(userId int, statusParam string) ([]domains.OrdersDomain, error) {
	tx, err := p.conn.Beginx()
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	query := `
		SELECT o.id, o.user_id, o.total_price, o.discounted_price, o.status, o.street, o.region, o.apartment, o.created_at, o.updated_at,
			u.id "user.id", u.name "user.name", u.phone "user.phone", r.name "user.role.name",
			u.country_id "user.country_id", u.street "user.street", u.region "user.region", u.apartment "user.apartment",
			u.date_of_birth "user.date_of_birth", u.created_at "user.created_at", u.updated_at "user.updated_at"
		FROM orders o
		JOIN users u ON o.user_id = u.id
		LEFT JOIN roles r ON u.role_id = r.id
		WHERE o.user_id = $1
	`
	args := []interface{}{userId}

	if statusParam != "" {
		query += " AND o.status = $2"
		args = append(args, statusParam)
	}

	var ordersRecord []records.Orders

	err = tx.Select(&ordersRecord, query, args...)
	if err != nil {
		tx.Rollback()
		return nil, helpers.PostgresErrorTransform(err)
	}

	queryOrderDetails := `
			SELECT od.id, od.order_id, od.product_id, od.quantity, od.price, od.sub_total,
				p.id "product.id", p.name "product.name", p.description "product.description", p.price "product.price", p.image "product.image", p.images "product.images", p.stock "product.stock", p.created_at "product.created_at", p.updated_at "product.updated_at",
				COALESCE(d.id, -1) "product.discount.id", COALESCE(d.product_id, 0) "product.discount.product_id", COALESCE(d.discount, 0) "product.discount.discount", COALESCE(NULLIF(d.start_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) "product.discount.start_date", COALESCE(NULLIF(d.end_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) "product.discount.end_date",
				CASE WHEN d.discount IS NOT NULL THEN p.price - (p.price * d.discount / 100) ELSE p.price END AS "product.discounted_price"
			FROM order_details od
			JOIN products p ON od.product_id = p.id
			LEFT JOIN discounts d ON p.id = d.product_id AND d.start_date <= NOW() AND d.end_date >= NOW()
			WHERE od.order_id = $1
		`

	for i, order := range ordersRecord {
		var orderDetailsRecord []records.OrderDetails
		err = tx.Select(&orderDetailsRecord, queryOrderDetails, order.Id)
		if err != nil {
			tx.Rollback()
			return nil, helpers.PostgresErrorTransform(err)
		}
		ordersRecord[i].OrderDetails = orderDetailsRecord
	}

	return records.ToArrayOfOrdersDomain(ordersRecord), nil
}

func (p *postgreOrdersRepository) Update(orders domains.OrdersDomain) error {
	query := `
		UPDATE orders
		SET status = COALESCE(:status, status),
		    region = COALESCE(:region, region),
		    street = COALESCE(:street, street), 
		    apartment = COALESCE(:apartment, apartment), 
		    updated_at = NOW()
		WHERE id = :id
		`

	var orderRecord records.Orders
	orderRecord = records.FromOrdersDomain(orders)

	_, err := p.conn.NamedExec(query, &orderRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreOrdersRepository) FindById(id int) (domains.OrdersDomain, error) {
	query := `SELECT * FROM orders WHERE id = $1`
	var orderRecord records.Orders
	err := p.conn.Get(&orderRecord, query, id)
	if err != nil {
		return domains.OrdersDomain{}, helpers.PostgresErrorTransform(err)
	}

	return orderRecord.ToDomain(), nil
}
