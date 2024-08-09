package requests

import "ga_marketplace/internal/business/domains"

type CreateSubserviceRequest struct {
	Name      string `json:"name" validate:"required"`
	ServiceId int    `json:"service_id" validate:"required"`
}

type UpdateSubserviceRequest struct {
	Name      *string `json:"name"`
	ServiceId *int    `json:"service_id"`
}

func (c *CreateSubserviceRequest) ToDomain() domains.SubServicesDomain {
	return domains.SubServicesDomain{
		Name:      c.Name,
		ServiceId: c.ServiceId,
	}
}
