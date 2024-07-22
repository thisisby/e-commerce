package requests

import "ga_marketplace/internal/business/domains"

type CreateOrderRequest struct {
	Street    string `json:"street" validate:"required"`
	Region    string `json:"region" validate:"required"`
	Apartment string `json:"apartment" validate:"required"`
	CityId    int    `json:"city_id" validate:"required"`
}

func (r *CreateOrderRequest) ToDomain() domains.OrdersDomain {
	return domains.OrdersDomain{
		Street:    r.Street,
		Region:    r.Region,
		Apartment: r.Apartment,
		CityId:    r.CityId,
	}
}

type UpdateOrderRequest struct {
	Status    *string `json:"status" validate:"orderstatus"`
	Street    *string `json:"street"`
	Region    *string `json:"region"`
	Apartment *string `json:"apartment"`
	CityId    *int    `json:"city_id"`
}

func (r *UpdateOrderRequest) ToDomain() *domains.OrdersDomain {
	return &domains.OrdersDomain{
		Status:    *r.Status,
		Street:    *r.Street,
		Region:    *r.Region,
		Apartment: *r.Apartment,
		CityId:    *r.CityId,
	}
}
