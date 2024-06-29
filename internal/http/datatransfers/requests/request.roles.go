package requests

import "ga_marketplace/internal/business/domains"

type CreateRoleRequest struct {
	Name string `json:"name" validate:"required"`
}

func (c *CreateRoleRequest) ToDomain() *domains.RoleDomain {
	return &domains.RoleDomain{
		Name: c.Name,
	}
}
