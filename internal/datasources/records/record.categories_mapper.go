package records

import "ga_marketplace/internal/business/domains"

func (r *Categories) ToDomain() domains.CategoriesDomain {
	return domains.CategoriesDomain{
		Id:   r.Id,
		Name: r.Name,
	}
}

func FromCategoriesDomain(domain domains.CategoriesDomain) Categories {
	return Categories{
		Id:   domain.Id,
		Name: domain.Name,
	}
}

func ToArrayOfCategoriesDomain(data []Categories) (result []domains.CategoriesDomain) {
	for _, v := range data {
		result = append(result, v.ToDomain())
	}
	return
}
