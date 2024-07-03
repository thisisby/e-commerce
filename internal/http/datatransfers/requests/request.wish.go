package requests

import "ga_marketplace/internal/business/domains"

type WishCreateRequest struct {
	ProductId int `json:"product_id" validate:"required"`
}

func (r *WishCreateRequest) ToDomain() *domains.WishDomain {
	return &domains.WishDomain{
		ProductId: r.ProductId,
	}
}
