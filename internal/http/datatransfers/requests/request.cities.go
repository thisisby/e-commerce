package requests

import "ga_marketplace/internal/business/domains"

type CreateCityRequest struct {
	Name                 string `json:"name" validate:"required"`
	DeliveryDurationDays int    `json:"delivery_duration_days" validate:"required"`
}

func (c *CreateCityRequest) ToDomain() domains.CityDomain {
	return domains.CityDomain{
		Name:                 c.Name,
		DeliveryDurationDays: c.DeliveryDurationDays,
	}
}

type UpdateCityRequest struct {
	Name                 string `json:"name" validate:"required"`
	DeliveryDurationDays int    `json:"delivery_duration_days" validate:"required"`
}

func (u *UpdateCityRequest) ToDomain() domains.CityDomain {
	return domains.CityDomain{
		Name:                 u.Name,
		DeliveryDurationDays: u.DeliveryDurationDays,
	}
}
