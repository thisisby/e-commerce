package records

import (
	"github.com/lib/pq"
	"time"
)

type Products struct {
	Id              int                  `db:"id"`
	Name            string               `db:"name"`
	Description     string               `db:"description"`
	Ingredients     string               `db:"ingredients"`
	Article         string               `db:"article"`
	CCode           string               `db:"c_code"`
	EdIzm           string               `db:"ed_izm"`
	Price           float64              `db:"price"`
	Weight          *float64             `db:"weight"`
	DiscountedPrice float64              `db:"discounted_price"`
	TotalPrice      *float64             `db:"total_price"`
	Discount        *Discounts           `db:"discount"`
	SubcategoryId   int                  `db:"subcategory_id"`
	Subcategory     *SubcategoriesRecord `db:"subcategory"`
	BrandId         int                  `db:"brand_id"`
	Brand           *Brands              `db:"brand"`
	Image           string               `db:"image"`
	Images          pq.StringArray       `db:"images"`
	IsInCart        bool                 `db:"is_in_cart"`
	IsInWishlist    bool                 `db:"is_in_wishlist"`
	CreatedAt       time.Time            `db:"created_at"`
	UpdatedAt       time.Time            `db:"updated_at"`
}
