package responses

import "ga_marketplace/internal/business/domains"

type CountryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func FromCountryDomain(inDom *domains.CountryDomain) CountryResponse {
	return CountryResponse{
		Id:   inDom.Id,
		Name: inDom.Name,
	}
}

func ToArrayOfCountryResponse(inDom []domains.CountryDomain) []CountryResponse {
	var outDom []CountryResponse

	for _, dom := range inDom {
		outDom = append(outDom, FromCountryDomain(&dom))
	}

	return outDom
}
