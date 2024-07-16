package records

import "ga_marketplace/internal/business/domains"

func (r *Countries) ToDomain() *domains.CountryDomain {
	return &domains.CountryDomain{
		Id:   r.Id,
		Name: r.Name,
	}
}

func FromCountryDomain(domain *domains.CountryDomain) *Countries {
	return &Countries{
		Id:   domain.Id,
		Name: domain.Name,
	}
}

func ToArrayOfCountryDomain(data []Countries) (result []domains.CountryDomain) {
	for _, v := range data {
		result = append(result, *v.ToDomain())
	}
	return
}
