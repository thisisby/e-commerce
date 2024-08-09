package records

import "ga_marketplace/internal/business/domains"

func (r *ServiceAddress) ToDomain() domains.ServiceAddressDomain {
	return domains.ServiceAddressDomain{
		Id:      r.Id,
		CityId:  r.CityId,
		City:    *r.City.ToDomain(),
		Address: r.Address,
	}
}

func FromServiceAddressDomain(domain domains.ServiceAddressDomain) ServiceAddress {
	return ServiceAddress{
		Id:      domain.Id,
		CityId:  domain.CityId,
		City:    *FromCityDomain(&domain.City),
		Address: domain.Address,
	}
}

func ToArrayOfServiceAddressDomain(data []ServiceAddress) (result []domains.ServiceAddressDomain) {
	for _, v := range data {
		result = append(result, v.ToDomain())
	}
	return
}
