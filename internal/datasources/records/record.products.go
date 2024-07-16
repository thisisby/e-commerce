package records

import (
	"github.com/lib/pq"
	"time"
)

type Products struct {
	Id              int            `db:"id"`
	Name            string         `db:"name"`
	Description     string         `db:"description"`
	Price           float64        `db:"price"`
	DiscountedPrice float64        `db:"discounted_price"`
	TotalPrice      *float64       `db:"total_price"`
	Discount        *Discounts     `db:"discount"`
	Image           string         `db:"image"`
	Images          pq.StringArray `db:"images"`
	Stock           int            `db:"stock"`
	IsInCart        bool           `db:"is_in_cart"`
	IsInWishlist    bool           `db:"is_in_wishlist"`
	CreatedAt       time.Time      `db:"created_at"`
	UpdatedAt       time.Time      `db:"updated_at"`
}
