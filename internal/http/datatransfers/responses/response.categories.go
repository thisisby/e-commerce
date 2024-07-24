package responses

import "ga_marketplace/internal/business/domains"

type CategoryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func FromCategoryDomain(inDom *domains.CategoriesDomain) *CategoryResponse {
	return &CategoryResponse{
		Id:   inDom.Id,
		Name: inDom.Name,
	}
}

func ToArrayOfCategoryResponse(inDom []domains.CategoriesDomain) []CategoryResponse {
	var outDom []CategoryResponse

	for _, dom := range inDom {
		outDom = append(outDom, *FromCategoryDomain(&dom))
	}

	return outDom
}
