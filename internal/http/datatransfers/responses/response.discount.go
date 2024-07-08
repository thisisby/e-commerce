package responses

import (
	"ga_marketplace/internal/business/domains"
	"time"
)

type DiscountResponse struct {
	Id        int       `json:"id"`
	ProductId int       `json:"product_id"`
	Discount  float64   `json:"discount"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func FromDiscountDomain(inDom *domains.DiscountsDomain) DiscountResponse {
	return DiscountResponse{
		Id:        inDom.Id,
		ProductId: inDom.ProductId,
		Discount:  inDom.Discount,
		StartDate: inDom.StartDate,
		EndDate:   inDom.EndDate,
	}
}
