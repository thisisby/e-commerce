package records

import "ga_marketplace/internal/business/domains"

func (r *SubcategoriesRecord) ToDomain() *domains.SubcategoriesDomain {
	if r == nil {
		return nil
	}
	return &domains.SubcategoriesDomain{
		Id:         r.Id,
		Name:       r.Name,
		CategoryId: r.CategoryId,
	}
}

func FromSubcategoriesDomain(domain *domains.SubcategoriesDomain) *SubcategoriesRecord {
	return &SubcategoriesRecord{
		Id:         domain.Id,
		Name:       domain.Name,
		CategoryId: domain.CategoryId,
	}
}

func ToArrayOfSubcategoriesDomain(data []SubcategoriesRecord) (result []domains.SubcategoriesDomain) {
	for _, v := range data {
		result = append(result, *v.ToDomain())
	}
	return
}
