package requests

import "ga_marketplace/internal/business/domains"

type CreateOrderRequest struct {
	Street               string               `json:"street" validate:"required"`
	Region               string               `json:"region" validate:"required"`
	Apartment            string               `json:"apartment" validate:"required"`
	CityId               int                  `json:"city_id" validate:"required"`
	Email                string               `json:"email" validate:"required,email"`
	StreetNum            string               `json:"street_num" validate:"required"`
	DeliveryMethod       string               `json:"delivery_method" validate:"order_delivery_method"`
	ReceiptUrl           string               `json:"receipt_url"`
	CreatePaymentRequest CreatePaymentRequest `json:"payment" validate:"required"`
}

func (r *CreateOrderRequest) ToDomain() domains.OrdersDomain {
	return domains.OrdersDomain{
		Street:         r.Street,
		Region:         r.Region,
		Apartment:      r.Apartment,
		CityId:         r.CityId,
		Email:          r.Email,
		StreetNum:      r.StreetNum,
		DeliveryMethod: r.DeliveryMethod,
		ReceiptUrl:     r.ReceiptUrl,
	}
}

type UpdateOrderRequest struct {
	Status         *string `json:"status" validate:"orderstatus"`
	Street         *string `json:"street"`
	Region         *string `json:"region"`
	Apartment      *string `json:"apartment"`
	CityId         *int    `json:"city_id"`
	StreetNum      *string `json:"street_num"`
	Email          *string `json:"email"`
	DeliveryMethod *string `json:"delivery_method" validate:"order_delivery_method"`
}

func (r *UpdateOrderRequest) ToDomain() *domains.OrdersDomain {
	return &domains.OrdersDomain{
		Status:         *r.Status,
		Street:         *r.Street,
		Region:         *r.Region,
		Apartment:      *r.Apartment,
		CityId:         *r.CityId,
		StreetNum:      *r.StreetNum,
		Email:          *r.Email,
		DeliveryMethod: *r.DeliveryMethod,
	}
}
