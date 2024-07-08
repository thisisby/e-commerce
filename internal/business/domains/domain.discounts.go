package domains

import "time"

type DiscountsDomain struct {
	Id        int
	ProductId int
	Discount  float64
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DiscountsRepository interface {
	Save(discount *DiscountsDomain) error
	FindByProductId(productId int) (*DiscountsDomain, error)
	DeleteByProductId(id int) error
}

type DiscountsUsecase interface {
	Save(discount *DiscountsDomain) (statusCode int, err error)
	DeleteByProductId(id int) (statusCode int, err error)
}
