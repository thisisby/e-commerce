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
	SubcategoryId   int
	Subcategory     *SubcategoriesDomain
	Image           string
	Images          []string
	Stock           int
	IsInCart        bool
	IsInWishlist    bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type ProductsRepository interface {
	FindById(id int) (*ProductDomain, error)
	Save(product *ProductDomain) error
	FindAllForMe(id int) ([]ProductDomain, error)
	UpdateById(inDom ProductDomain) error
	FindAllForMeBySubcategoryId(id int, subcategoryId int) ([]ProductDomain, error)
}

type ProductsUsecase interface {
	Save(product *ProductDomain) (int, error)
	FindAllForMe(id int) ([]ProductDomain, int, error)
	UpdateById(inDom ProductDomain) (int, error)
	FindById(id int) (*ProductDomain, int, error)
	FindAllForMeBySubcategoryId(id int, subcategoryId int) ([]ProductDomain, int, error)
}
