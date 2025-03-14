package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/constants"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type postgreOrdersRepository struct {
	conn *sqlx.DB
}

func NewPostgreOrdersRepository(conn *sqlx.DB) domains.OrdersRepository {
	return &postgreOrdersRepository{
		conn: conn,
	}
}

func (p *postgreOrdersRepository) Save(orders domains.OrdersDomain) (int, error) {
	tx, err := p.conn.Begin()
	if err != nil {
		return 0, helpers.PostgresErrorTransform(err)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	query := `
		INSERT INTO orders (user_id, total_price, discounted_price, city_id, status, street, region, apartment, street_num, email, delivery_method, receipt_url)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id
	`
	var orderId int
	err = tx.QueryRow(query, orders.UserId, orders.TotalPrice, orders.DiscountedPrice, orders.CityId, orders.Status, orders.Street, orders.Region, orders.Apartment, orders.StreetNum, orders.Email, orders.DeliveryMethod, orders.ReceiptUrl).Scan(&orderId)
	if err != nil {
		tx.Rollback()
		return 0, helpers.PostgresErrorTransform(err)
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
			return 0, helpers.PostgresErrorTransform(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return orderId, nil
}

func (p *postgreOrdersRepository) FindByUserId(userId int, statusParam string) ([]domains.OrdersDomain, error) {
	tx, err := p.conn.Beginx()
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	query := `
		SELECT 
		    o.id, o.user_id, o.total_price, o.discounted_price,
		    o.status, o.street, o.region, o.apartment, o.street_num,
		    o.email, o.delivery_method, o.created_at, o.updated_at, o.receipt_url,
			u.id "user.id", u.name "user.name", u.phone "user.phone",
			r.name "user.role.name", u.city_id "user.city_id", u.street "user.street",
			u.region "user.region", u.apartment "user.apartment",
			u.date_of_birth "user.date_of_birth", u.email "user.email",
			u.street_num "user.street_num", u.created_at "user.created_at",
			u.updated_at "user.updated_at", c.id "city.id", c.name "city.name",
			c.delivery_duration_days "city.delivery_duration_days"
		FROM orders o
		JOIN users u ON o.user_id = u.id
		LEFT JOIN roles r ON u.role_id = r.id
		LEFT JOIN cities c ON o.city_id = c.id
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
			SELECT 
			    od.id, od.order_id, od.product_id, od.quantity, od.price, od.sub_total,
			    p.id "product.id", p.name "product.name", p.description "product.description",
			    p.ingredients "product.ingredients", p.c_code "product.c_code", p.ed_izm "product.ed_izm",
			    p.article "product.article", p.subcategory_id "product.subcategory_id", p.brand_id "product.brand_id",
				p.price "product.price", p.weight "product.weight", p.image "product.image", p.images "product.images",
				p.created_at "product.created_at", p.updated_at "product.updated_at",
				s.id "product.subcategory.id", s.name "product.subcategory.name",
				s.category_id "product.subcategory.category_id", b.id "product.brand.id",
				b.name "product.brand.name",
				COALESCE(d.id, -1) "product.discount.id", COALESCE(d.product_id, 0) "product.discount.product_id",
				COALESCE(d.discount, 0) "product.discount.discount", 
				COALESCE(NULLIF(d.start_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) "product.discount.start_date",
				COALESCE(NULLIF(d.end_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) "product.discount.end_date",
				CASE WHEN d.discount IS NOT NULL THEN p.price - (p.price * d.discount / 100) ELSE p.price END AS "product.discounted_price",
				COALESCE(SUM(CASE 
                    WHEN psi.transaction_type = 1 THEN psi.quantity 
                    WHEN psi.transaction_type = 2 THEN -psi.quantity 
                    ELSE 0 
                 END), 0) AS "product.stock"
			FROM order_details od
			JOIN products p ON od.product_id = p.id
			LEFT JOIN discounts d ON p.id = d.product_id AND d.start_date <= NOW() AND d.end_date >= NOW()
			JOIN subcategories s ON p.subcategory_id = s.id
			JOIN brands b ON p.brand_id = b.id
			LEFT JOIN product_stock_item psi ON p.c_code = psi.product_code
			LEFT JOIN product_stock ps ON psi.transaction_id = ps.transaction_id AND ps.active = TRUE
			WHERE od.order_id = $1
			GROUP BY od.id, p.id, d.id, s.id, b.id
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

	mappedOrders := records.ToArrayOfOrdersDomain(ordersRecord)

	return mappedOrders, nil
}

func (p *postgreOrdersRepository) Update(orders domains.OrdersDomain) error {
	query := `
		UPDATE orders
		SET status = COALESCE(:status, status),
		    region = COALESCE(:region, region),
		    street = COALESCE(:street, street), 
		    apartment = COALESCE(:apartment, apartment), 
		    street_num = COALESCE(:street_num, street_num),
		    email = COALESCE(:email, email),
		    city_id = COALESCE(:city_id, city_id),
		    delivery_method = COALESCE(:delivery_method, delivery_method),
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

func (p *postgreOrdersRepository) FindAll(filter constants.OrderFilter) ([]domains.OrdersDomain, error) {
	tx, err := p.conn.Beginx()
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	query := `
		SELECT o.id, o.user_id, o.total_price, o.discounted_price, o.status, o.street, o.region, o.apartment, o.email, o.street_num, o.delivery_method, o.created_at, o.updated_at, o.receipt_url,
			u.id "user.id", u.name "user.name", u.phone "user.phone", r.name "user.role.name",
			u.city_id "user.city_id", u.street "user.street", u.region "user.region", u.apartment "user.apartment",
			u.email "user.email", u.street_num "user.street_num",
			u.date_of_birth "user.date_of_birth", u.created_at "user.created_at", u.updated_at "user.updated_at",
			c.id "city.id", c.name "city.name", c.delivery_duration_days "city.delivery_duration_days"
		FROM orders o
		JOIN users u ON o.user_id = u.id
		LEFT JOIN roles r ON u.role_id = r.id
		LEFT JOIN cities c ON o.city_id = c.id
		WHERE 1 = 1
	`
	var args []interface{}
	argIndex := 1

	if filter.Status != nil && *filter.Status != "" {
		query += " AND o.status = $" + strconv.Itoa(argIndex)
		args = append(args, *filter.Status)
		argIndex++
	}
	if filter.Limit != nil {
		query += " LIMIT $" + strconv.Itoa(argIndex)
		args = append(args, *filter.Limit)
		argIndex++
	}
	if filter.Offset != nil {
		query += " OFFSET $" + strconv.Itoa(argIndex)
		args = append(args, *filter.Offset)
		argIndex++
	}

	var ordersRecord []records.Orders

	err = tx.Select(&ordersRecord, query, args...)
	if err != nil {
		tx.Rollback()
		return nil, helpers.PostgresErrorTransform(err)
	}

	queryOrderDetails := `
			SELECT 
			    od.id, od.order_id, od.product_id, od.quantity, od.price, od.sub_total,
			    p.id "product.id", p.name "product.name", p.description "product.description",
			    p.ingredients "product.ingredients", p.c_code "product.c_code", p.ed_izm "product.ed_izm",
			    p.article "product.article", p.weight "product.weight", p.subcategory_id "product.subcategory_id", p.brand_id "product.brand_id",
				p.price "product.price", p.image "product.image", p.images "product.images",
				p.created_at "product.created_at", p.updated_at "product.updated_at",
				s.id "product.subcategory.id", s.name "product.subcategory.name", s.category_id "product.subcategory.category_id",
				b.id "product.brand.id", b.name "product.brand.name",
				COALESCE(d.id, -1) "product.discount.id", COALESCE(d.product_id, 0) "product.discount.product_id",
				COALESCE(d.discount, 0) "product.discount.discount", 
				COALESCE(NULLIF(d.start_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) "product.discount.start_date",
				COALESCE(NULLIF(d.end_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) "product.discount.end_date",
				CASE WHEN d.discount IS NOT NULL THEN p.price - (p.price * d.discount / 100) ELSE p.price END AS "product.discounted_price",
				COALESCE(SUM(CASE 
                    WHEN psi.transaction_type = 1 THEN psi.quantity 
                    WHEN psi.transaction_type = 2 THEN -psi.quantity 
                    ELSE 0 
                 END), 0) AS "product.stock"
			FROM order_details od
			JOIN products p ON od.product_id = p.id
			LEFT JOIN discounts d ON p.id = d.product_id AND d.start_date <= NOW() AND d.end_date >= NOW()
			JOIN subcategories s ON p.subcategory_id = s.id
			JOIN brands b ON p.brand_id = b.id
			LEFT JOIN product_stock_item psi ON p.c_code = psi.product_code
			LEFT JOIN product_stock ps ON psi.transaction_id = ps.transaction_id AND ps.active = TRUE
			WHERE od.order_id = $1
			GROUP BY od.id, p.id, d.id, s.id, b.id
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

func (p *postgreOrdersRepository) Cancel(id int) error {
	query := `UPDATE orders SET status = 'canceled' WHERE id = $1`
	_, err := p.conn.Exec(query, id)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}
