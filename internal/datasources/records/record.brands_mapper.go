package records

import "ga_marketplace/internal/business/domains"

func (b *Brands) ToDomain() *domains.BrandsDomain {
	if b == nil {
		return nil
	}
	return &domains.BrandsDomain{
		Id:   b.Id,
		Name: b.Name,
		Info: b.Info,
	}
}

func FromBrandsDomain(domain domains.BrandsDomain) Brands {
	return Brands{
		Id:   domain.Id,
		Name: domain.Name,
		Info: domain.Info,
	}
}

func ToArrayOfBrandsDomain(data []Brands) (result []domains.BrandsDomain) {
	for _, v := range data {
		result = append(result, *v.ToDomain())
	}
	return
}
