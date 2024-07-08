package records

import "time"

type Discounts struct {
	Id        int       `db:"id"`
	ProductId int       `db:"product_id"`
	Discount  float64   `db:"discount"`
	StartDate time.Time `db:"start_date"`
	EndDate   time.Time `db:"end_date"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
