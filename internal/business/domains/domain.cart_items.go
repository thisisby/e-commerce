package domains

import "time"

type CartItemsDomain struct {
	Id        int
	UserId    int
	ProductId int
	Product   ProductDomain
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartItemsRepository interface {
	FindAllByUserId(id int) ([]CartItemsDomain, error)
	FindByUserIdAndProductId(userId int, productId int) (*CartItemsDomain, error)
	FindById(id int) (*CartItemsDomain, error)
	Save(cart *CartItemsDomain) error
	DeleteByIdAndUserId(id int, userId int) error
	UpdateByIdAndUserId(cart *CartItemsDomain) error
}

type CartUsecase interface {
	FindAllByUserId(id int) (outDom []CartItemsDomain, statusCode int, err error)
	Save(inDom *CartItemsDomain) (statusCode int, err error)
	DeleteByIdAndUserId(id int, userId int) (statusCode int, err error)
	UpdateByIdAndUserId(id int, userId int, cart *CartItemsDomain) (statusCode int, err error)
}
