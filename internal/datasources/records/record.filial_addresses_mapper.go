package records

import (
	"ga_marketplace/internal/business/domains"
)

func (f *FilialAddresses) ToDomain() domains.FilialAddressesDomain {
	return domains.FilialAddressesDomain{
		Id:        f.Id,
		Street:    f.Street,
		Region:    f.Region,
		Apartment: f.Apartment,
		StreetNum: f.StreetNum,
		CityId:    f.CityId,
		City:      f.City.ToDomain(),
	}
}

func ToArrayOfFilialAddressesDomain(fs []FilialAddresses) []domains.FilialAddressesDomain {
	var res []domains.FilialAddressesDomain
	for _, v := range fs {
		res = append(res, v.ToDomain())
	}

	return res
}

func FromFilialAddressesDomain(d domains.FilialAddressesDomain) FilialAddresses {
	return FilialAddresses{
		Id:        d.Id,
		Street:    d.Street,
		Region:    d.Region,
		Apartment: d.Apartment,
		StreetNum: d.StreetNum,
		CityId:    d.CityId,
	}
}
