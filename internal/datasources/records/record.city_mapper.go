package records

import "ga_marketplace/internal/business/domains"

func (r *Cities) ToDomain() *domains.CityDomain {
	if r == nil || r.Id == 0 {
		return nil
	}
	return &domains.CityDomain{
		Id:                   r.Id,
		Name:                 r.Name,
		DeliveryDurationDays: r.DeliveryDurationDays,
	}
}

func FromCityDomain(domain *domains.CityDomain) *Cities {
	if domain == nil {
		return nil
	}
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
