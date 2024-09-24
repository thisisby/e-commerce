package requests

import "ga_marketplace/internal/business/domains"

type CreateBrandRequest struct {
	Name string `json:"name" validate:"required"`
	Info string `json:"info" validate:"required"`
}

func (c *CreateBrandRequest) ToDomain() domains.BrandsDomain {
	return domains.BrandsDomain{
		Name: c.Name,
		Info: c.Info,
	}
}

type UpdateBrandRequest struct {
	Name string `json:"name" validate:"required"`
	Info string `json:"info" validate:"required"`
}

func (u *UpdateBrandRequest) ToDomain() domains.BrandsDomain {
	return domains.BrandsDomain{
		Name: u.Name,
		Info: u.Info,
	}
}
