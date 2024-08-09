package responses

import "ga_marketplace/internal/business/domains"

type ServiceResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func FromServiceDomain(inDom *domains.ServicesDomain) *ServiceResponse {
	return &ServiceResponse{
		Id:   inDom.Id,
		Name: inDom.Name,
	}
}

func ToArrayOfServiceResponse(inDom []domains.ServicesDomain) []ServiceResponse {
	var outDom []ServiceResponse

	for _, dom := range inDom {
		outDom = append(outDom, *FromServiceDomain(&dom))
	}

	return outDom
}
