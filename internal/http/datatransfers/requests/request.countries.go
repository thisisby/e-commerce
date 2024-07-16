package requests

import "ga_marketplace/internal/business/domains"

type CreateCountryRequest struct {
	Name string `json:"name"`
}

func (c *CreateCountryRequest) ToDomain() domains.CountryDomain {
	return domains.CountryDomain{
		Name: c.Name,
	}
}

type UpdateCountryRequest struct {
	Name string `json:"name"`
}

func (u *UpdateCountryRequest) ToDomain() domains.CountryDomain {
	return domains.CountryDomain{
		Name: u.Name,
	}
}
