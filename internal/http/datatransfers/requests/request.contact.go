package requests

import "ga_marketplace/internal/business/domains"

type CreateContactRequest struct {
	Title string `json:"title" validate:"required"`
	Value string `json:"value" validate:"required"`
}

func (c *CreateContactRequest) ToDomain() domains.ContactDomain {
	return domains.ContactDomain{
		Title: c.Title,
		Value: c.Value,
	}
}

type UpdateContactRequest struct {
	Title *string `json:"title"`
	Value *string `json:"value"`
}
