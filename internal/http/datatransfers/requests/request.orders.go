package requests

import "ga_marketplace/internal/business/domains"

type CreateOrderRequest struct {
	Street    string `json:"street"`
	Region    string `json:"region"`
	Apartment string `json:"apartment"`
}

func (r *CreateOrderRequest) ToDomain() domains.OrdersDomain {
	return domains.OrdersDomain{
		Street:    r.Street,
		Region:    r.Region,
		Apartment: r.Apartment,
	}
}

type UpdateOrderRequest struct {
	Status    *string `json:"status" validate:"orderstatus"`
	Street    *string `json:"street"`
	Region    *string `json:"region"`
	Apartment *string `json:"apartment"`
}

func (r *UpdateOrderRequest) ToDomain() *domains.OrdersDomain {
	return &domains.OrdersDomain{
		Status:    *r.Status,
		Street:    *r.Street,
		Region:    *r.Region,
		Apartment: *r.Apartment,
	}
}
