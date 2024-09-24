package postgre

import (
	"fmt"
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log/slog"
	"strconv"
	"strings"
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
		SELECT 
		    p.id, p.name, p.description, p.ingredients, p.c_code, 
		    p.ed_izm, p.article, p.price, p.subcategory_id, p.brand_id, 
		    p.image, p.images, p.weight, p.country_of_production, p.sex,
		    p.volume, p.created_at, p.updated_at,
		    s.id "subcategory.id", s.name "subcategory.name", 
		    s.category_id "subcategory.category_id",
		    b.id "brand.id", b.name "brand.name",
		    COALESCE(d.id, -1) "discount.id", COALESCE(d.product_id, 0) "discount.product_id", 
		    COALESCE(d.discount, 0) "discount.discount", 
		    COALESCE(NULLIF(d.start_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) "discount.start_date", 
		    COALESCE(NULLIF(d.end_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) "discount.end_date",
		    CASE WHEN d.discount IS NOT NULL THEN p.price - (p.price * d.discount / 100) ELSE p.price END AS "discounted_price",
		     COALESCE(SUM(CASE 
                    WHEN psi.transaction_type = 1 AND ps.active = TRUE THEN psi.quantity 
                    WHEN psi.transaction_type = 2 AND ps.active = TRUE THEN -psi.quantity 
                    ELSE 0 
                 END), 0) AS "stock"
		FROM products p
		LEFT JOIN discounts d ON p.id = d.product_id AND d.start_date <= NOW() AND d.end_date >= NOW()
		LEFT JOIN product_stock_item psi ON p.c_code = psi.product_code
		LEFT JOIN product_stock ps ON psi.transaction_id = ps.transaction_id
		JOIN subcategories s ON p.subcategory_id = s.id
		JOIN brands b ON p.brand_id = b.id
		WHERE p.id = $1
		GROUP BY p.id, d.id, s.id, b.id
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
		INSERT INTO products (name, description, price, subcategory_id, image, images, brand_id)
		VALUES (:name, :description, :price, :subcategory_id, :image, :images, :brand_id)		
		`

	productRecord := records.FromProductDomain(product)

	_, err := p.conn.NamedQuery(query, productRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreProductsRepository) FindAllForMe(id int, filter domains.ProductFilter) ([]domains.ProductDomain, error) {
	query := `
		SELECT 
    		p.id, p.name, p.description, p.price, p.ingredients, p.c_code, 
    		p.ed_izm, 
    p.article, 
    p.image, 
    p.images, 
    p.weight,
    p.created_at, 
    p.updated_at, 
    p.subcategory_id, 
    p.brand_id,
    COALESCE(d.id, -1) AS "discount.id", 
    COALESCE(d.product_id, 0) AS "discount.product_id", 
    COALESCE(d.discount, 0) AS "discount.discount", 
    COALESCE(NULLIF(d.start_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) AS "discount.start_date", 
    COALESCE(NULLIF(d.end_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) AS "discount.end_date",
    s.id AS "subcategory.id", 
    s.name AS "subcategory.name", 
    s.category_id AS "subcategory.category_id",
    b.id AS "brand.id", 
    b.name AS "brand.name",
    CASE 
        WHEN c.product_id IS NOT NULL THEN TRUE 
        ELSE FALSE 
    END AS "is_in_cart",
    CASE 
        WHEN w.product_id IS NOT NULL THEN w.id 
        ELSE -1 
    END AS "is_in_wishlist",
    CASE 
        WHEN d.discount IS NOT NULL THEN p.price - (p.price * d.discount / 100) 
        ELSE p.price 
    END AS "discounted_price",
    COALESCE(SUM(CASE 
                    WHEN psi.transaction_type = 1 AND ps.active = TRUE THEN psi.quantity 
                    WHEN psi.transaction_type = 2 AND ps.active = TRUE THEN -psi.quantity 
                    ELSE 0 
                 END), 0) AS "stock"
FROM 
    products p
LEFT JOIN 
    cart_items c ON p.id = c.product_id AND c.user_id = $1
LEFT JOIN 
    wishes w ON p.id = w.product_id AND w.user_id = $1
LEFT JOIN 
    discounts d ON p.id = d.product_id AND d.start_date <= NOW() AND d.end_date >= NOW()
LEFT JOIN 
    product_stock_item psi ON p.c_code = psi.product_code
LEFT JOIN 
    product_stock ps ON psi.transaction_id = ps.transaction_id
LEFT JOIN 
    subcategories s ON p.subcategory_id = s.id
LEFT JOIN 
    brands b ON p.brand_id = b.id
		`

	var filters []string
	var args []interface{}
	args = append(args, id)

	if filter.Name != "" {
		filters = append(filters, "p.name ILIKE '%' || $"+strconv.Itoa(len(args)+1)+" || '%'")
		args = append(args, filter.Name)
	}
	if filter.MinPrice != "" {
		filters = append(filters, "p.price >= $"+strconv.Itoa(len(args)+1))
		args = append(args, filter.MinPrice)
	}
	if filter.MaxPrice != "" {
		filters = append(filters, "p.price <= $"+strconv.Itoa(len(args)+1))
		args = append(args, filter.MaxPrice)
	}
	if filter.SubcategoryID != "" {
		filters = append(filters, "p.subcategory_id = $"+strconv.Itoa(len(args)+1))
		args = append(args, filter.SubcategoryID)
	}
	if filter.BrandID != "" {
		filters = append(filters, "p.brand_id = $"+strconv.Itoa(len(args)+1))
		args = append(args, filter.BrandID)
	}
	if filter.CountryOfProduction != "" {
		filters = append(filters, "p.country_of_production = $"+strconv.Itoa(len(args)+1))
		args = append(args, filter.CountryOfProduction)
	}
	if filter.Volume != 0 {
		filters = append(filters, "p.volume = $"+strconv.Itoa(len(args)+1))
		args = append(args, filter.Volume)
	}
	if filter.Sex != "" {
		filters = append(filters, "p.sex = $"+strconv.Itoa(len(args)+1))
		args = append(args, filter.Sex)
	}

	if len(filter.Attributes) > 1 {
		filters = append(filters, "a.name = ANY($"+strconv.Itoa(len(args)+1)+")")
		args = append(args, pq.Array(filter.Attributes))

		query += `
		JOIN product_attributes pa ON p.id = pa.product_id
		JOIN attributes a ON pa.attribute_id = a.id
		`
	}

	if len(filters) > 0 {
		query += " WHERE " + strings.Join(filters, " AND ")
	}

	query += " GROUP BY p.id, s.id, b.id, d.id, c.id, w.id"

	offset := (filter.Page - 1) * filter.PageSize
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", filter.PageSize, offset)

	var productsRecord []records.Products

	err := p.conn.Select(&productsRecord, query, args...)
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
		SET 
		    weight = $1,
		    image = $2,
		    images = $3,
		    subcategory_id = $4,
		    brand_id = $5,
			country_of_production = $6,
			volume = $7,
			sex = $8
		WHERE id = $9
	`

	_, err = tx.Exec(updateQuery,
		productRecord.Weight,
		productRecord.Image,
		productRecord.Images,
		productRecord.SubcategoryId,
		productRecord.BrandId,
		productRecord.CountryOfProduction,
		productRecord.Volume,
		productRecord.Sex,
		productRecord.Id,
	)
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

func (p *postgreProductsRepository) FindAllForMeBySubcategoryId(id int, subcategoryId int) ([]domains.ProductDomain, error) {
	query := `
		SELECT 
		    p.id, p.name, p.description, p.price, p.image, 
		    p.ingredients, p.c_code, p.ed_izm, p.article, p.weight,
		    p.images, p.created_at, p.updated_at, p.subcategory_id, p.brand_id,
			COALESCE(d.id, -1) "discount.id", COALESCE(d.product_id, 0) "discount.product_id", 
			COALESCE(d.discount, 0) "discount.discount", 
			COALESCE(NULLIF(d.start_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) "discount.start_date", 
			COALESCE(NULLIF(d.end_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) "discount.end_date",
			s.id "subcategory.id", s.name "subcategory.name", s.category_id "subcategory.category_id",
			b.id "brand.id", b.name "brand.name",
		CASE WHEN c.product_id IS NOT NULL THEN TRUE ELSE FALSE END AS "is_in_cart",
		CASE WHEN w.product_id IS NOT NULL THEN w.id ELSE -1 END AS "is_in_wishlist",
		CASE WHEN d.discount IS NOT NULL THEN p.price - (p.price * d.discount / 100) ELSE p.price END AS "discounted_price",
		COALESCE(SUM(CASE 
                    WHEN psi.transaction_type = 1 AND ps.active = TRUE THEN psi.quantity 
                    WHEN psi.transaction_type = 2 AND ps.active = TRUE THEN -psi.quantity 
                    ELSE 0 
                 END), 0) AS "stock"
		FROM products p
		LEFT JOIN cart_items c ON p.id = c.product_id AND c.user_id = $1
		LEFT JOIN wishes w ON p.id = w.product_id AND w.user_id = $1
		LEFT JOIN discounts d ON p.id = d.product_id AND d.start_date <= NOW() AND d.end_date >= NOW()
		LEFT JOIN subcategories s ON p.subcategory_id = s.id
		LEFT JOIN product_stock_item psi ON p.c_code = psi.product_code
		LEFT JOIN product_stock ps ON psi.transaction_id = ps.transaction_id
		JOIN brands b ON p.brand_id = b.id
		WHERE p.subcategory_id = $2
		GROUP BY p.id, d.id, s.id, b.id, c.product_id, w.id
		`

	var productsRecord []records.Products

	err := p.conn.Select(&productsRecord, query, id, subcategoryId)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfProductsDomain(productsRecord), nil
}

func (p *postgreProductsRepository) FindAllForMeByBrandId(id int, brandId int) ([]domains.ProductDomain, error) {
	query := `
		SELECT
		    p.id, p.name, p.description, p.price, p.image, p.images, 
		    p.ingredients, p.c_code, p.ed_izm, p.article, p.weight,
		    p.created_at, p.updated_at, p.subcategory_id, p.brand_id,
			COALESCE(d.id, -1) "discount.id", COALESCE(d.product_id, 0) "discount.product_id", 
			COALESCE(d.discount, 0) "discount.discount",
			COALESCE(NULLIF(d.start_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) "discount.start_date",
			COALESCE(NULLIF(d.end_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) "discount.end_date",
			s.id "subcategory.id", s.name "subcategory.name", s.category_id "subcategory.category_id",
			b.id "brand.id", b.name "brand.name",
		CASE WHEN c.product_id IS NOT NULL THEN TRUE ELSE FALSE END AS "is_in_cart",
		CASE WHEN w.product_id IS NOT NULL THEN w.id ELSE -1 END AS "is_in_wishlist",
		CASE WHEN d.discount IS NOT NULL THEN p.price - (p.price * d.discount / 100) ELSE p.price END AS "discounted_price",
		COALESCE(SUM(CASE 
                    WHEN psi.transaction_type = 1 AND ps.active = TRUE THEN psi.quantity 
                    WHEN psi.transaction_type = 2 AND ps.active = TRUE THEN -psi.quantity 
                    ELSE 0 
                 END), 0) AS "stock"
		FROM products p
		LEFT JOIN cart_items c ON p.id = c.product_id AND c.user_id = $1
		LEFT JOIN wishes w ON p.id = w.product_id AND w.user_id = $1
		LEFT JOIN discounts d ON p.id = d.product_id AND d.start_date <= NOW() AND d.end_date >= NOW()
		LEFT JOIN product_stock_item psi ON p.c_code = psi.product_code
		LEFT JOIN product_stock ps ON psi.transaction_id = ps.transaction_id
		JOIN subcategories s ON p.subcategory_id = s.id
		JOIN brands b ON p.brand_id = b.id
		WHERE p.brand_id = $2
		GROUP BY p.id, d.id, s.id, b.id, c.product_id, w.id
		`

	var productsRecord []records.Products

	err := p.conn.Select(&productsRecord, query, id, brandId)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfProductsDomain(productsRecord), nil
}

func (p *postgreProductsRepository) SaveFrom1c(product *domains.ProductDomainV2) error {
	query := `
		INSERT INTO products (
		                      name, 
		                      description, 
		                      price, 
		                      article,
		                      c_code,
		                      ed_izm,
		                      subcategory_id, 
		                      brand_id,
		                      ingredients,
		                      image, 
		                      created_at,
		                      updated_at
		                      )
		VALUES (
		        :name, 
		        :description, 
		        :price, 
		        :article, 
		        :c_code, 
		        :ed_izm,
		        :subcategory_id, 
		        :brand_id,
		        :ingredients,
		        :image, 
		        NOW(),
		        NOW()
		        )
		`

	_, err := p.conn.NamedQuery(query, product)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreProductsRepository) UpdateFrom1c(code string, product *domains.ProductDomain) error {
	query := `
		UPDATE products
		SET 
		    name = $1,
		    description = $2,
		    price = $3,
		    article = $4,
		    c_code = $5,
		    ed_izm = $6,
		    updated_at = NOW()
		WHERE c_code = $7
		`

	_, err := p.conn.Exec(query,
		product.Name,
		product.Description,
		product.Price,
		product.Article,
		product.CCode,
		product.EdIzm,
		code,
	)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreProductsRepository) FindByCode(code string) (*domains.ProductDomain, error) {
	query := `
		SELECT 
		    p.id, p.name, p.description, p.ingredients, p.c_code, 
		    p.ed_izm, p.article, p.price, p.subcategory_id, p.brand_id, 
		    p.image, p.images, p.weight, p.created_at, p.updated_at,
		    s.id "subcategory.id", s.name "subcategory.name", 
		    s.category_id "subcategory.category_id",
		    b.id "brand.id", b.name "brand.name",
		    COALESCE(d.id, -1) "discount.id", COALESCE(d.product_id, 0) "discount.product_id", 
		    COALESCE(d.discount, 0) "discount.discount", 
		    COALESCE(NULLIF(d.start_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) "discount.start_date", 
		    COALESCE(NULLIF(d.end_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) "discount.end_date",
		    CASE WHEN d.discount IS NOT NULL THEN p.price - (p.price * d.discount / 100) ELSE p.price END AS "discounted_price",
		    COALESCE(SUM(CASE 
                    WHEN psi.transaction_type = 1 AND ps.active = TRUE THEN psi.quantity 
                    WHEN psi.transaction_type = 2 AND ps.active = TRUE THEN -psi.quantity 
                    ELSE 0 
                 END), 0) AS "stock"
		FROM products p
		LEFT JOIN discounts d ON p.id = d.product_id AND d.start_date <= NOW() AND d.end_date >= NOW()
		LEFT JOIN product_stock_item psi ON p.c_code = psi.product_code
		LEFT JOIN product_stock ps ON psi.transaction_id = ps.transaction_id
		JOIN subcategories s ON p.subcategory_id = s.id
		JOIN brands b ON p.brand_id = b.id
		WHERE p.c_code = $1
		GROUP BY p.id, d.id, s.id, b.id
		`

	var productRecord records.Products

	err := p.conn.Get(&productRecord, query, code)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return productRecord.ToDomain(), nil
}

func (p *postgreProductsRepository) FindByIdForUser(id int, userId int) (*domains.ProductDomain, error) {
	query := `
		SELECT 
		    p.id, p.name, p.description, p.ingredients, p.c_code, 
		    p.ed_izm, p.article, p.price, p.subcategory_id, p.brand_id, 
		    p.image, p.images, p.weight, p.country_of_production, p.sex,
		    p.volume, p.created_at, p.updated_at,
		    s.id "subcategory.id", s.name "subcategory.name", 
		    s.category_id "subcategory.category_id",
		    b.id "brand.id", b.name "brand.name", b.info "brand.info",
		    ch.name "characteristic",
		    COALESCE(d.id, -1) "discount.id", COALESCE(d.product_id, 0) "discount.product_id", 
		    COALESCE(d.discount, 0) "discount.discount", 
		    COALESCE(NULLIF(d.start_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) "discount.start_date", 
		    COALESCE(NULLIF(d.end_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) "discount.end_date",
		    CASE WHEN d.discount IS NOT NULL THEN p.price - (p.price * d.discount / 100) ELSE p.price END AS "discounted_price",
		    CASE WHEN c.product_id IS NOT NULL THEN TRUE ELSE FALSE END AS "is_in_cart",
		    CASE WHEN w.product_id IS NOT NULL THEN w.id ELSE -1 END AS "is_in_wishlist",
		    COALESCE(SUM(CASE 
                    WHEN psi.transaction_type = 1 AND ps.active = TRUE THEN psi.quantity 
                    WHEN psi.transaction_type = 2 AND ps.active = TRUE THEN -psi.quantity 
                    ELSE 0 
                 END), 0) AS "stock",
			array_agg(COALESCE(a.name, '')) AS "attributes"
		FROM products p
		LEFT JOIN discounts d ON p.id = d.product_id AND d.start_date <= NOW() AND d.end_date >= NOW()
		LEFT JOIN cart_items c ON p.id = c.product_id AND c.user_id = $2
		LEFT JOIN wishes w ON p.id = w.product_id AND w.user_id = $2
		LEFT JOIN product_stock_item psi ON p.c_code = psi.product_code
		LEFT JOIN product_stock ps ON psi.transaction_id = ps.transaction_id
		JOIN subcategories s ON p.subcategory_id = s.id
		JOIN brands b ON p.brand_id = b.id
		LEFT JOIN characteristics ch ON ch.subcategory_id = p.subcategory_id
		LEFT JOIN product_attributes pa ON p.id = pa.product_id
		LEFT JOIN attributes a ON pa.attribute_id = a.id
		WHERE p.id = $1
		GROUP BY p.id, s.id, b.id, d.id, c.id, w.id, ps.transaction_id, ch.name
		`

	var productRecord records.Products

	err := p.conn.Get(&productRecord, query, id, userId)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return productRecord.ToDomain(), nil
}

func (p *postgreProductsRepository) FindAll(filter domains.ProductFilter) ([]domains.ProductDomain, int, error) {
	query := `
		SELECT 
			p.id, p.name, p.description, p.price, p.ingredients, p.c_code, 
			p.ed_izm, p.article, p.image, p.images, p.weight, p.created_at, 
	p.updated_at, 	p.subcategory_id,	p.brand_id,	COALESCE(d.id, -1) AS "discount.id", 
	COALESCE(d.product_id, 0) AS "discount.product_id",	COALESCE(d.discount, 0) AS "discount.discount", 
	COALESCE(NULLIF(d.start_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) AS "discount.start_date", 
	COALESCE(NULLIF(d.end_date, '0001-01-01'::timestamp), '1970-01-01'::timestamp) AS "discount.end_date",
	s.id AS "subcategory.id", 	s.name AS "subcategory.name",	s.category_id AS "subcategory.category_id",
	b.id AS "brand.id", 	b.name AS "brand.name", 	
	FALSE AS "is_in_cart",
	-1 AS "is_in_wishlist",
	CASE 
		WHEN d.discount IS NOT NULL THEN p.price - (p.price * d.discount / 100) 
		ELSE p.price 
	END AS "discounted_price",
	COALESCE(SUM(CASE 
					WHEN psi.transaction_type = 1 AND ps.active = TRUE THEN psi.quantity 
					WHEN psi.transaction_type = 2 AND ps.active = TRUE THEN -psi.quantity 
					ELSE 0 
				 END), 0) AS "stock"
	FROM
		products p
	LEFT JOIN 
		discounts d ON p.id = d.product_id AND d.start_date <= NOW() AND d.end_date >= NOW()
	LEFT JOIN
		product_stock_item psi ON p.c_code = psi.product_code
	LEFT JOIN
		product_stock ps ON psi.transaction_id = ps.transaction_id
	JOIN 
		subcategories s ON p.subcategory_id = s.id
	JOIN
		brands b ON p.brand_id = b.id
	`

	var filters []string
	var args []interface{}

	if filter.Name != "" {
		filters = append(filters, "p.name ILIKE '%' || $"+strconv.Itoa(len(args)+1)+" || '%'")
		args = append(args, filter.Name)
	}

	if filter.MinPrice != "" {
		filters = append(filters, "p.price >= $"+strconv.Itoa(len(args)+1))
		args = append(args, filter.MinPrice)
	}

	if filter.MaxPrice != "" {
		filters = append(filters, "p.price <= $"+strconv.Itoa(len(args)+1))
		args = append(args, filter.MaxPrice)
	}

	if filter.SubcategoryID != "" {
		filters = append(filters, "p.subcategory_id = $"+strconv.Itoa(len(args)+1))
		args = append(args, filter.SubcategoryID)
	}

	if filter.BrandID != "" {
		filters = append(filters, "p.brand_id = $"+strconv.Itoa(len(args)+1))
		args = append(args, filter.BrandID)
	}
	if filter.CountryOfProduction != "" {
		filters = append(filters, "p.country_of_production = $"+strconv.Itoa(len(args)+1))
		args = append(args, filter.CountryOfProduction)
	}
	if filter.Volume != 0 {
		filters = append(filters, "p.volume = $"+strconv.Itoa(len(args)+1))
		args = append(args, filter.Volume)
	}
	if filter.Sex != "" {
		filters = append(filters, "p.sex = $"+strconv.Itoa(len(args)+1))
		args = append(args, filter.Sex)
	}

	if len(filter.Attributes) > 1 {
		filters = append(filters, "a.name = ANY($"+strconv.Itoa(len(args)+1)+")")
		args = append(args, pq.Array(filter.Attributes))

		query += `
		JOIN product_attributes pa ON p.id = pa.product_id
		JOIN attributes a ON pa.attribute_id = a.id
		`
	}

	if len(filters) > 0 {
		query += " WHERE " + strings.Join(filters, " AND ")
	}

	query += " GROUP BY p.id, d.id, s.id, b.id"

	offset := (filter.Page - 1) * filter.PageSize
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", filter.PageSize, offset)

	var productsRecord []records.Products

	err := p.conn.Select(&productsRecord, query, args...)
	if err != nil {
		return nil, 0, helpers.PostgresErrorTransform(err)
	}

	slog.Info("query", query)

	return records.ToArrayOfProductsDomain(productsRecord), len(productsRecord), nil
}

func (p *postgreProductsRepository) AddAttributesToProduct(productId int, attributes []int) error {
	tx, err := p.conn.Begin()
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, attributeId := range attributes {
		slog.Info("attributeId", attributeId)
		query := `
			INSERT INTO product_attributes (product_id, attribute_id)
			VALUES ($1, $2)
			`

		_, err = tx.Exec(query, productId, attributeId)
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

func (p *postgreProductsRepository) DeleteAttributesFromProduct(productId int, attributeIds []int) error {
	tx, err := p.conn.Begin()
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, attributeId := range attributeIds {
		query := `
			DELETE FROM product_attributes
			WHERE product_id = $1 AND attribute_id = $2
			`

		_, err = tx.Exec(query, productId, attributeId)
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
