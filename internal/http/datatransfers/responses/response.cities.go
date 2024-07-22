package responses

import "ga_marketplace/internal/business/domains"

type CityResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func FromCityDomain(inDom *domains.CityDomain) *CityResponse {
	return &CityResponse{
		Id:   inDom.Id,
		Name: inDom.Name,
	}
}

func ToArrayOfCityResponse(inDom []domains.CityDomain) []CityResponse {
	var outDom []CityResponse

	for _, dom := range inDom {
		outDom = append(outDom, *FromCityDomain(&dom))
	}

	return outDom
}
