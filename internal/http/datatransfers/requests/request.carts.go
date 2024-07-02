package requests

import "ga_marketplace/internal/business/domains"

type CartCreateRequest struct {
	ProductId int `json:"product_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required"`
}

func (r *CartCreateRequest) ToDomain() *domains.CartItemsDomain {
	return &domains.CartItemsDomain{
		ProductId: r.ProductId,
		Quantity:  r.Quantity,
	}
}
