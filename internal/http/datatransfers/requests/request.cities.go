package requests

import "ga_marketplace/internal/business/domains"

type CreateCityRequest struct {
	Name string `json:"name"`
}

func (c *CreateCityRequest) ToDomain() domains.CityDomain {
	return domains.CityDomain{
		Name: c.Name,
	}
}

type UpdateCityRequest struct {
	Name string `json:"name"`
}

func (u *UpdateCityRequest) ToDomain() domains.CityDomain {
	return domains.CityDomain{
		Name: u.Name,
	}
}
