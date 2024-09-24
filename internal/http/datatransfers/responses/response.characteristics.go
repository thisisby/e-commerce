package responses

import "ga_marketplace/internal/business/domains"

type CharacteristicsResponse struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	SubcategoryId int    `json:"subcategory_id"`
}

func FromCharacteristicDomain(domain domains.CharacteristicsDomain) CharacteristicsResponse {
	return CharacteristicsResponse{
		Id:            domain.Id,
		Name:          domain.Name,
		SubcategoryId: domain.SubcategoryId,
	}
}

func ToArrayOfCharacteristicsResponse(data []domains.CharacteristicsDomain) (result []CharacteristicsResponse) {
	for _, v := range data {
		result = append(result, FromCharacteristicDomain(v))
	}
	return
}
