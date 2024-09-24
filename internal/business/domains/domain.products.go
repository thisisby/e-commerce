package domains

import (
	"time"
)

type ProductDomain struct {
	Id                  int
	Name                string
	Description         string
	Ingredients         string
	Article             string
	CCode               string
	EdIzm               string
	Price               float64
	Weight              *float64
	DiscountedPrice     float64
	TotalPrice          *float64
	Discount            *DiscountsDomain
	SubcategoryId       int
	Subcategory         *SubcategoriesDomain
	BrandId             int
	Brand               *BrandsDomain
	Image               string
	Images              []string
	IsInCart            bool
	IsInWishlist        int
	Stock               int
	Attributes          []string
	Characteristic      string
	CountryOfProduction string
	Volume              float64
	Sex                 string
	CreatedAt           time.Time
	UpdatedAt           time.Time
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

type ProductFilter struct {
	Name                string
	MinPrice            string
	MaxPrice            string
	SubcategoryID       string
	BrandID             string
	Attributes          []string
	CountryOfProduction string
	Volume              float64
	Sex                 string
	Page                int
	PageSize            int
}

type ProductPagination struct {
	Page     int
	PageSize int
}

type ProductsRepository interface {
	FindById(id int) (*ProductDomain, error)
	FindByIdForUser(id int, userId int) (*ProductDomain, error)
	Save(product *ProductDomain) error
	SaveFrom1c(product *ProductDomainV2) error
	FindAllForMe(id int, filter ProductFilter) ([]ProductDomain, error)
	UpdateById(inDom ProductDomain) error
	FindAllForMeBySubcategoryId(id int, subcategoryId int) ([]ProductDomain, error)
	FindAllForMeByBrandId(id int, brandId int) ([]ProductDomain, error)
	UpdateFrom1c(code string, product *ProductDomain) error
	FindByCode(code string) (*ProductDomain, error)
	FindAll(filter ProductFilter) ([]ProductDomain, int, error)
	AddAttributesToProduct(productId int, attributes []int) error
	DeleteAttributesFromProduct(productId int, attributeIds []int) error
	DeleteById(id int) error
}

type ProductsUsecase interface {
	Save(product *ProductDomain) (int, error)
	SaveFrom1c(product *ProductDomainV2) (int, error)
	FindAllForMe(id int, filter ProductFilter) ([]ProductDomain, int, error)
	UpdateById(inDom ProductDomain) (int, error)
	FindById(id int) (*ProductDomain, int, error)
	FindAllForMeBySubcategoryId(id int, subcategoryId int) ([]ProductDomain, int, error)
	FindAllForMeByBrandId(id int, brandId int) ([]ProductDomain, int, error)
	UpdateFrom1c(code string, product *ProductDomain) (int, error)
	FindByCode(code string) (*ProductDomain, int, error)
	FindByIdForUser(id int, userId int) (*ProductDomain, int, error)
	FindAll(filter ProductFilter) ([]ProductDomain, int, error)
	AddAttributesToProduct(productId int, attributes []int) (int, error)
	DeleteAttributesFromProduct(productId int, attributeIds []int) (int, error)
	DeleteById(id int) (int, error)
}
