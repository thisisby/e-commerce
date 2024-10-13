package responses

import "ga_marketplace/internal/business/domains"

type FilialAddressesResponse struct {
	Id        int           `json:"id"`
	Street    string        `json:"street"`
	Region    string        `json:"region"`
	Apartment string        `json:"apartment"`
	StreetNum string        `json:"street_num"`
	CityId    int           `json:"city_id"`
	City      *CityResponse `json:"city"`
}

func FromFilialAddressesDomain(domain domains.FilialAddressesDomain) FilialAddressesResponse {
	return FilialAddressesResponse{
		Id:        domain.Id,
		Street:    domain.Street,
		Region:    domain.Region,
		Apartment: domain.Apartment,
		StreetNum: domain.StreetNum,
		CityId:    domain.CityId,
		City:      FromCityDomain(domain.City),
	}
}

func ToArrayOfFilialAddressesResponse(filialAddresses []domains.FilialAddressesDomain) []FilialAddressesResponse {
	var result []FilialAddressesResponse
	for _, v := range filialAddresses {
		result = append(result, FromFilialAddressesDomain(v))
	}
	return result
}
