package requests

import "ga_marketplace/internal/business/domains"

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

func (c *CreateCategoryRequest) ToDomain() domains.CategoriesDomain {
	return domains.CategoriesDomain{
		Name: c.Name,
	}
}

type UpdateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

func (u *UpdateCategoryRequest) ToDomain() domains.CategoriesDomain {
	return domains.CategoriesDomain{
		Name: u.Name,
	}
}
