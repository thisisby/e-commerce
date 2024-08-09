package requests

import "ga_marketplace/internal/business/domains"

type CreateServiceRequest struct {
	Name string `json:"name" validate:"required"`
}

func (c *CreateServiceRequest) ToDomain() domains.ServicesDomain {
	return domains.ServicesDomain{
		Name: c.Name,
	}
}

type UpdateServiceRequest struct {
	Name string `json:"name" validate:"required"`
}

func (u *UpdateServiceRequest) ToDomain() domains.ServicesDomain {
	return domains.ServicesDomain{
		Name: u.Name,
	}
}
