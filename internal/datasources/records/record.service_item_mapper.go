package records

import "ga_marketplace/internal/business/domains"

func (r *ServiceItem) ToDomain() domains.ServiceItemDomain {
	return domains.ServiceItemDomain{
		Id:           r.Id,
		Title:        r.Title,
		Duration:     r.Duration,
		Description:  r.Description,
		Price:        r.Price,
		SubServiceId: r.SubServiceId,
	}
}

func FromServiceItemDomain(domain domains.ServiceItemDomain) ServiceItem {
	return ServiceItem{
		Id:           domain.Id,
		Title:        domain.Title,
		Duration:     domain.Duration,
		Description:  domain.Description,
		Price:        domain.Price,
		SubServiceId: domain.SubServiceId,
	}
}

func ToArrayOfServiceItemDomain(data []ServiceItem) (result []domains.ServiceItemDomain) {
	for _, v := range data {
		result = append(result, v.ToDomain())
	}
	return
}
