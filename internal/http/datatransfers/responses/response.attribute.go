package responses

import "ga_marketplace/internal/business/domains"

type AttributesResponse struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	CharacteristicsId int    `json:"characteristic_id"`
}

func FromAttributeDomain(domain domains.AttributesDomain) AttributesResponse {
	return AttributesResponse{
		Id:                domain.Id,
		Name:              domain.Name,
		CharacteristicsId: domain.CharacteristicsId,
	}
}

func ToArrayOfAttributesResponse(data []domains.AttributesDomain) (result []AttributesResponse) {
	for _, v := range data {
		result = append(result, FromAttributeDomain(v))
	}
	return
}
