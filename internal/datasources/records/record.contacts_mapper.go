package records

import "ga_marketplace/internal/business/domains"

func (r *Contacts) ToDomain() *domains.ContactDomain {
	return &domains.ContactDomain{
		Id:    r.Id,
		Title: r.Title,
		Value: r.Value,
	}
}

func FromContactDomain(domain *domains.ContactDomain) *Contacts {
	return &Contacts{
		Id:    domain.Id,
		Title: domain.Title,
		Value: domain.Value,
	}
}

func ToArrayOfContactDomain(data []Contacts) (result []domains.ContactDomain) {
	for _, v := range data {
		result = append(result, *v.ToDomain())
	}
	return
}
