package responses

import "ga_marketplace/internal/business/domains"

type ServiceAddressResponse struct {
	Id      int          `json:"id"`
	CityId  int          `json:"city_id"`
	City    CityResponse `json:"city"`
	Address string       `json:"address"`
}

func FromServiceAddressDomain(domain domains.ServiceAddressDomain) ServiceAddressResponse {
	return ServiceAddressResponse{
		Id:     domain.Id,
		CityId: domain.CityId,
		City: CityResponse{
			Id:   domain.City.Id,
			Name: domain.City.Name,
		},
		Address: domain.Address,
	}
}

func ToArrayOfServiceAddressResponse(data []domains.ServiceAddressDomain) (result []ServiceAddressResponse) {
	for _, v := range data {
		result = append(result, FromServiceAddressDomain(v))
	}
	return
}
