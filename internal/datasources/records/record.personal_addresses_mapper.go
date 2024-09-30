package records

import "ga_marketplace/internal/business/domains"

func (p *PersonalAddresses) ToDomain() domains.PersonalAddressesDomain {
	return domains.PersonalAddressesDomain{
		Id:        p.Id,
		UserId:    p.UserId,
		Street:    p.Street,
		Region:    p.Region,
		Apartment: p.Apartment,
		StreetNum: p.StreetNum,
		CityId:    p.CityId,
		City:      p.City.ToDomain(),
		User:      p.User.ToDomain(),
	}
}

func ToArrayOfPersonalAddressesDomain(personalAddresses []PersonalAddresses) []domains.PersonalAddressesDomain {
	var result []domains.PersonalAddressesDomain
	for _, v := range personalAddresses {
		result = append(result, v.ToDomain())
	}
	return result
}

func FromPersonalAddressesDomain(domain domains.PersonalAddressesDomain) PersonalAddresses {
	return PersonalAddresses{
		Id:        domain.Id,
		UserId:    domain.UserId,
		Street:    domain.Street,
		Region:    domain.Region,
		Apartment: domain.Apartment,
		StreetNum: domain.StreetNum,
		CityId:    domain.CityId,
	}
}
