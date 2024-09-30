package responses

import "ga_marketplace/internal/business/domains"

type PersonalAddressesResponse struct {
	Id        int           `json:"id"`
	UserId    int           `json:"user_id"`
	User      *UserResponse `json:"user"`
	Street    string        `json:"street"`
	Region    string        `json:"region"`
	Apartment string        `json:"apartment"`
	StreetNum string        `json:"street_num"`
	CityId    int           `json:"city_id"`
	City      *CityResponse `json:"city"`
}

func FromPersonalAddressesDomain(domain domains.PersonalAddressesDomain) PersonalAddressesResponse {
	return PersonalAddressesResponse{
		Id:        domain.Id,
		UserId:    domain.UserId,
		Street:    domain.Street,
		Region:    domain.Region,
		Apartment: domain.Apartment,
		StreetNum: domain.StreetNum,
		CityId:    domain.CityId,
		City:      FromCityDomain(domain.City),
	}
}

func ToArrayOfPersonalAddressesResponse(personalAddresses []domains.PersonalAddressesDomain) []PersonalAddressesResponse {
	var result []PersonalAddressesResponse
	for _, v := range personalAddresses {
		result = append(result, FromPersonalAddressesDomain(v))
	}
	return result
}
