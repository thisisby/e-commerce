package records

import (
	"ga_marketplace/internal/business/domains"
	"time"
)

type Orders struct {
	Id              int            `db:"id"`
	UserId          int            `db:"user_id"`
	User            Users          `db:"user"`
	OrderDetails    []OrderDetails `db:"order_details"`
	TotalPrice      float64        `db:"total_price"`
	DiscountedPrice float64        `db:"discounted_price"`
	Status          string         `db:"status"`
	Street          string         `db:"street"`
	Region          string         `db:"region"`
	Apartment       string         `db:"apartment"`
	CreatedAt       time.Time      `db:"created_at"`
	UpdatedAt       time.Time      `db:"updated_at"`
}

type OrderDetails struct {
	Id        int      `db:"id"`
	OrderId   int      `db:"order_id"`
	ProductId int      `db:"product_id"`
	Product   Products `db:"product"`
	Quantity  int      `db:"quantity"`
	Price     float64  `db:"price"`
	SubTotal  float64  `db:"sub_total"`
}

func (r *OrderDetails) ToDomain() domains.OrderDetailsDomain {
	domain := domains.OrderDetailsDomain{
		Id:        r.Id,
		OrderId:   r.OrderId,
		ProductId: r.ProductId,
		Product:   *r.Product.ToDomain(),
		Quantity:  r.Quantity,
		Price:     r.Price,
		SubTotal:  r.SubTotal,
	}

	return domain
}

func ToArrayOfOrderDetailsDomain(data []OrderDetails) []domains.OrderDetailsDomain {
	var result []domains.OrderDetailsDomain
	for _, val := range data {
		result = append(result, val.ToDomain())
	}

	return result
}

func (r *Orders) ToDomain() domains.OrdersDomain {
	domain := domains.OrdersDomain{
		Id:              r.Id,
		UserId:          r.UserId,
		User:            *r.User.ToDomain(),
		OrderDetails:    ToArrayOfOrderDetailsDomain(r.OrderDetails),
		TotalPrice:      r.TotalPrice,
		DiscountedPrice: r.DiscountedPrice,
		Status:          r.Status,
		Street:          r.Street,
		Region:          r.Region,
		Apartment:       r.Apartment,
	}

	return domain
}
