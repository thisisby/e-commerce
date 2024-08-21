package domains

import "time"

type ProductDomain struct {
	Id              int
	Name            string
	Description     string
	Ingredients     string
	Article         string
	CCode           string
	EdIzm           string
	Price           float64
	DiscountedPrice float64
	TotalPrice      *float64
	Discount        *DiscountsDomain
	SubcategoryId   int
	Subcategory     *SubcategoriesDomain
	BrandId         int
	Brand           *BrandsDomain
	Image           string
	Images          []string
	IsInCart        bool
	IsInWishlist    bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type ProductDomainV2 struct {
	Id            int     `db:"id" json:"id"`
	Name          string  `db:"name" json:"name"`
	Description   string  `db:"description" json:"description"`
	Price         float64 `db:"price" json:"price"`
	Article       string  `db:"article" json:"article"`
	CCode         string  `db:"c_code" json:"c_code"`
	EdIzm         string  `db:"ed_izm" json:"ed_izm"`
	Ingredients   string  `db:"ingredients" json:"ingredients"`
	BrandId       int     `db:"brand_id" json:"brand_id"`
	SubcategoryId int     `db:"subcategory_id" json:"subcategory_id"`
	Image         string  `db:"image" json:"image"`
}

type ProductsRepository interface {
	FindById(id int) (*ProductDomain, error)
	Save(product *ProductDomain) error
	SaveFrom1c(product *ProductDomainV2) error
	FindAllForMe(id int) ([]ProductDomain, error)
	UpdateById(inDom ProductDomain) error
	FindAllForMeBySubcategoryId(id int, subcategoryId int) ([]ProductDomain, error)
	FindAllForMeByBrandId(id int, brandId int) ([]ProductDomain, error)
}

type ProductsUsecase interface {
	Save(product *ProductDomain) (int, error)
	SaveFrom1c(product *ProductDomainV2) (int, error)
	FindAllForMe(id int) ([]ProductDomain, int, error)
	UpdateById(inDom ProductDomain) (int, error)
	FindById(id int) (*ProductDomain, int, error)
	FindAllForMeBySubcategoryId(id int, subcategoryId int) ([]ProductDomain, int, error)
	FindAllForMeByBrandId(id int, brandId int) ([]ProductDomain, int, error)
}
