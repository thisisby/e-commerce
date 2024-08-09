package requests

import "ga_marketplace/internal/business/domains"

type ServiceItemCreateRequest struct {
	Title         string  `json:"title" validate:"required"`
	Duration      int     `json:"duration" validate:"required"`
	Description   string  `json:"description" validate:"required"`
	Price         float64 `json:"price" validate:"required"`
	Subservice_id int     `json:"subservice_id" validate:"required"`
}

type ServiceItemUpdateRequest struct {
	Title         *string  `json:"title"`
	Duration      *int     `json:"duration"`
	Description   *string  `json:"description"`
	Price         *float64 `json:"price"`
	Subservice_id *int     `json:"subservice_id"`
}

func (c *ServiceItemCreateRequest) ToDomain() domains.ServiceItemDomain {
	return domains.ServiceItemDomain{
		Title:        c.Title,
		Duration:     c.Duration,
		Description:  c.Description,
		Price:        c.Price,
		SubServiceId: c.Subservice_id,
	}
}
