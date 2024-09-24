package requests

import "ga_marketplace/internal/business/domains"

type CreateAttributeRequest struct {
	Name             string `json:"name" validate:"required"`
	CharacteristicId int    `json:"characteristic_id" validate:"required"`
}

func (c *CreateAttributeRequest) ToDomain() domains.AttributesDomain {
	return domains.AttributesDomain{
		Name:              c.Name,
		CharacteristicsId: c.CharacteristicId,
	}
}

type UpdateAttributeRequest struct {
	Name             *string `json:"name"`
	CharacteristicId *int    `json:"characteristic_id"`
}
