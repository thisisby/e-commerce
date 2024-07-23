package responses

import "ga_marketplace/internal/business/domains"

type ContactResponse struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Value string `json:"value"`
}

func FromContactDomain(inDom *domains.ContactDomain) ContactResponse {
	return ContactResponse{
		Id:    inDom.Id,
		Title: inDom.Title,
		Value: inDom.Value,
	}
}

func ToArrayOfContactResponse(inDom []domains.ContactDomain) []ContactResponse {
	var outDom []ContactResponse

	for _, dom := range inDom {
		outDom = append(outDom, FromContactDomain(&dom))
	}

	return outDom
}
