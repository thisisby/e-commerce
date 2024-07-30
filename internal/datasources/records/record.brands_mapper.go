package records

import "ga_marketplace/internal/business/domains"

func (b *Brands) ToDomain() *domains.BrandsDomain {
	return &domains.BrandsDomain{
		Id:   b.Id,
		Name: b.Name,
	}
}

func FromBrandsDomain(domain domains.BrandsDomain) Brands {
	return Brands{
		Id:   domain.Id,
		Name: domain.Name,
	}
}

func ToArrayOfBrandsDomain(data []Brands) (result []domains.BrandsDomain) {
	for _, v := range data {
		result = append(result, *v.ToDomain())
	}
	return
}
