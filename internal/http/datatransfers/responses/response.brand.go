package responses

import "ga_marketplace/internal/business/domains"

type BrandResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func FromBrandDomain(inDom *domains.BrandsDomain) *BrandResponse {
	return &BrandResponse{
		Id:   inDom.Id,
		Name: inDom.Name,
	}
}

func ToArrayOfBrandResponse(inDom []domains.BrandsDomain) []BrandResponse {
	var outDom []BrandResponse

	for _, dom := range inDom {
		outDom = append(outDom, *FromBrandDomain(&dom))
	}

	return outDom
}
