package requests

import "ga_marketplace/internal/business/domains"

type CreateSubcategoryRequest struct {
	Name       string `json:"name" validate:"required"`
	CategoryId int    `json:"category_id" validate:"required"`
}

type UpdateSubcategoryRequest struct {
	Name       *string `json:"name"`
	CategoryId *int    `json:"category_id"`
}

func (c *CreateSubcategoryRequest) ToDomain() domains.SubcategoriesDomain {
	return domains.SubcategoriesDomain{
		Name:       c.Name,
		CategoryId: c.CategoryId,
	}
}
