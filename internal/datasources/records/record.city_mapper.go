package records

import "ga_marketplace/internal/business/domains"

func (r *Cities) ToDomain() *domains.CityDomain {
	return &domains.CityDomain{
		Id:                   r.Id,
		Name:                 r.Name,
		DeliveryDurationDays: r.DeliveryDurationDays,
	}
}

func FromCityDomain(domain *domains.CityDomain) *Cities {
	return &Cities{
		Id:                   domain.Id,
		Name:                 domain.Name,
		DeliveryDurationDays: domain.DeliveryDurationDays,
	}
}

func ToArrayOfCityDomain(data []Cities) (result []domains.CityDomain) {
	for _, v := range data {
		result = append(result, *v.ToDomain())
	}
	return
}
