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
	FindByUserId(id int) ([]CartItemsDomain, error)
	FindByUserIdAndProductId(userId int, productId int) (*CartItemsDomain, error)
	FindById(id int) (*CartItemsDomain, error)
	Save(cart *CartItemsDomain) error
	Delete(id int, userId int) error
	FindAll() ([]CartItemsDomain, error)
	Update(id int, userId int, cart *CartItemsDomain) error
}

type CartUsecase interface {
	FindByUserId(id int) (outDom []CartItemsDomain, statusCode int, err error)
	Save(inDom *CartItemsDomain) (statusCode int, err error)
	Delete(id int, userId int) (statusCode int, err error)
	FindAll(userId int, isAdmin bool) (outDom []CartItemsDomain, statusCode int, err error)
	Update(id int, userId int, cart *CartItemsDomain) (statusCode int, err error)
}
