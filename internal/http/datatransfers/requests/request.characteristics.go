package requests

import "ga_marketplace/internal/business/domains"

type CreateCharacteristicRequest struct {
	Name          string `json:"name" validate:"required"`
	SubcategoryId int    `json:"subcategory_id" validate:"required"`
}

func (c *CreateCharacteristicRequest) ToDomain() domains.CharacteristicsDomain {
	return domains.CharacteristicsDomain{
		Name:          c.Name,
		SubcategoryId: c.SubcategoryId,
	}
}

type UpdateCharacteristicRequest struct {
	Name          *string `json:"name"`
	SubcategoryId *int    `json:"subcategory_id"`
}
