package requests

import "ga_marketplace/internal/business/domains"

type CreatePersonalAddressRequest struct {
	Street    string `json:"street" validate:"required"`
	Region    string `json:"region" validate:"required"`
	Apartment string `json:"apartment" validate:"required"`
	StreetNum string `json:"street_num" validate:"required"`
	CityId    int    `json:"city_id" validate:"required"`
}

func (c *CreatePersonalAddressRequest) ToDomain() domains.PersonalAddressesDomain {
	return domains.PersonalAddressesDomain{
		Street:    c.Street,
		Region:    c.Region,
		Apartment: c.Apartment,
		StreetNum: c.StreetNum,
		CityId:    c.CityId,
	}
}

type UpdatePersonalAddressRequest struct {
	Street    string `json:"street" validate:"required"`
	Region    string `json:"region" validate:"required"`
	Apartment string `json:"apartment" validate:"required"`
	StreetNum string `json:"street_num" validate:"required"`
	CityId    int    `json:"city_id" validate:"required"`
}

func (u *UpdatePersonalAddressRequest) ToDomain() domains.PersonalAddressesDomain {
	return domains.PersonalAddressesDomain{
		Street:    u.Street,
		Region:    u.Region,
		Apartment: u.Apartment,
		StreetNum: u.StreetNum,
		CityId:    u.CityId,
	}
}
