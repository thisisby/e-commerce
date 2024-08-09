package records

import "ga_marketplace/internal/business/domains"

func (r *Services) ToDomain() domains.ServicesDomain {
	return domains.ServicesDomain{
		Id:   r.Id,
		Name: r.Name,
	}
}

func FromServicesDomain(domain domains.ServicesDomain) Services {
	return Services{
		Id:   domain.Id,
		Name: domain.Name,
	}
}

func ToArrayOfServicesDomain(data []Services) (result []domains.ServicesDomain) {
	for _, v := range data {
		result = append(result, v.ToDomain())
	}
	return
}
