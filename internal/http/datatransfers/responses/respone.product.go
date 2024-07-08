package responses

import (
	"ga_marketplace/internal/business/domains"
	"time"
)

type ProductResponse struct {
	Id              int              `json:"id"`
	Name            string           `json:"name"`
	Description     string           `json:"description"`
	Price           float64          `json:"price"`
	DiscountedPrice float64          `json:"discounted_price"`
	TotalPrice      float64          `json:"total_price"`
	Discount        DiscountResponse `json:"discount"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
}

func FromProductDomain(inDom *domains.ProductDomain) ProductResponse {
	return ProductResponse{
		Id:              inDom.Id,
		Name:            inDom.Name,
		Description:     inDom.Description,
		Price:           inDom.Price,
		DiscountedPrice: inDom.DiscountedPrice,
		TotalPrice:      inDom.TotalPrice,
		Discount:        FromDiscountDomain(&inDom.Discount),
		CreatedAt:       inDom.CreatedAt,
		UpdatedAt:       inDom.UpdatedAt,
	}
}
