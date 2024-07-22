package domains

type OrdersDomain struct {
	Id              int
	UserId          int
	User            UserDomain
	OrderDetails    []OrderDetailsDomain
	TotalPrice      float64
	DiscountedPrice float64
	CityId          int
	City            CityDomain
	Status          string
	Street          string
	Region          string
	Apartment       string
}

type OrderDetailsDomain struct {
	Id        int
	OrderId   int
	ProductId int
	Product   ProductDomain
	Quantity  int
	Price     float64
	SubTotal  float64
}

type OrdersRepository interface {
	Save(orders OrdersDomain) error
	FindByUserId(userId int, statusParam string) ([]OrdersDomain, error)
	Update(orders OrdersDomain) error
	FindById(id int) (OrdersDomain, error)
}

type OrdersUsecase interface {
	Save(orders OrdersDomain, cartItems []CartItemsDomain, totalAmount CartItemTotalAmount) (int, error)
	FindByUserId(userId int, statusParam string) ([]OrdersDomain, int, error)
	Update(orders OrdersDomain) (int, error)
	FindById(id int) (OrdersDomain, int, error)
}
