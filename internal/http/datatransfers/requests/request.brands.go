package requests

import "ga_marketplace/internal/business/domains"

type CreateBrandRequest struct {
	Name string `json:"name" validate:"required"`
}

func (c *CreateBrandRequest) ToDomain() domains.BrandsDomain {
	return domains.BrandsDomain{
		Name: c.Name,
	}
}

type UpdateBrandRequest struct {
	Name string `json:"name" validate:"required"`
}

func (u *UpdateBrandRequest) ToDomain() domains.BrandsDomain {
	return domains.BrandsDomain{
		Name: u.Name,
	}
}
