package responses

import "ga_marketplace/internal/business/domains"

type SubcategoryResponse struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	CategoryId int    `json:"category_id"`
}

func FromSubcategoryDomain(inDom *domains.SubcategoriesDomain) *SubcategoryResponse {
	return &SubcategoryResponse{
		Id:         inDom.Id,
		Name:       inDom.Name,
		CategoryId: inDom.CategoryId,
	}
}

func ToArrayOfSubcategoryResponse(inDom []domains.SubcategoriesDomain) []SubcategoryResponse {
	var outDom []SubcategoryResponse

	for _, dom := range inDom {
		outDom = append(outDom, *FromSubcategoryDomain(&dom))
	}

	return outDom
}
