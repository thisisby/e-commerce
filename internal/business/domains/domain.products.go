package domains

import "time"

type ProductDomain struct {
	Id              int
	Name            string
	Description     string
	Price           float64
	DiscountedPrice float64
	TotalPrice      float64
	Discount        DiscountsDomain
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type ProductRepository interface {
	FindById(id int) (*ProductDomain, error)
}
