package records

import "ga_marketplace/internal/business/domains"

func (r *SubServiceRecord) ToDomain() *domains.SubServicesDomain {
	return &domains.SubServicesDomain{
		Id:        r.Id,
		Name:      r.Name,
		ServiceId: r.ServiceId,
	}
}

func FromSubServicesDomain(domain *domains.SubServicesDomain) *SubServiceRecord {
	return &SubServiceRecord{
		Id:        domain.Id,
		Name:      domain.Name,
		ServiceId: domain.ServiceId,
	}
}

func ToArrayOfSubServicesDomain(data []SubServiceRecord) (result []domains.SubServicesDomain) {
	for _, v := range data {
		result = append(result, *v.ToDomain())
	}
	return
}
