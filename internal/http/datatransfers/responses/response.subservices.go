package responses

import "ga_marketplace/internal/business/domains"

type SubServiceResponse struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ServiceId int    `json:"service_id"`
}

func FromSubServiceDomain(inDom *domains.SubServicesDomain) *SubServiceResponse {
	return &SubServiceResponse{
		Id:        inDom.Id,
		Name:      inDom.Name,
		ServiceId: inDom.ServiceId,
	}
}

func ToArrayOfSubServiceResponse(inDom []domains.SubServicesDomain) []SubServiceResponse {
	var outDom []SubServiceResponse

	for _, dom := range inDom {
		outDom = append(outDom, *FromSubServiceDomain(&dom))
	}

	return outDom
}
