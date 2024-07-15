package domains

import "time"

type ProductDomain struct {
	Id              int
	Name            string
	Description     string
	Price           float64
	DiscountedPrice float64
	TotalPrice      *float64
	Discount        *DiscountsDomain
	Image           string
	Images          []string
	IsInCart        bool
	IsInWishlist    bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type ProductsRepository interface {
	FindById(id int) (*ProductDomain, error)
	Save(product *ProductDomain) error
	FindAllForMe(id int) ([]ProductDomain, error)
}

type ProductsUsecase interface {
	Save(product *ProductDomain) (int, error)
	FindAllForMe(id int) ([]ProductDomain, int, error)
}
