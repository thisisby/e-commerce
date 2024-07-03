package domains

import "time"

type WishDomain struct {
	Id        int
	UserId    int
	ProductId int
	Product   ProductDomain
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WishRepository interface {
	FindByUserId(id int) ([]WishDomain, error)
	FindByUserIdAndProductId(userId int, productId int) (*WishDomain, error)
	FindById(id int) (*WishDomain, error)
	Save(wish *WishDomain) error
	Delete(id int, userId int) error
}

type WishUsecase interface {
	FindByUserId(id int) (outDom []WishDomain, statusCode int, err error)
	Save(inDom *WishDomain) (statusCode int, err error)
	Delete(id int, userId int) (statusCode int, err error)
}
