package requests

import (
	"ga_marketplace/internal/business/domains"
	"time"
)

type DiscountCreateRequest struct {
	ProductId int       `json:"product_id" validate:"required"`
	Discount  float64   `json:"discount" validate:"required"`
	StartDate time.Time `json:"start_date" validate:"required"`
	EndDate   time.Time `json:"end_date" validate:"required"`
}

func (d *DiscountCreateRequest) ToDomain() *domains.DiscountsDomain {
	return &domains.DiscountsDomain{
		ProductId: d.ProductId,
		Discount:  d.Discount,
		StartDate: d.StartDate,
		EndDate:   d.EndDate,
	}
}
