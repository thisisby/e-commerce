package records

import "ga_marketplace/internal/business/domains"

type Characteristics struct {
	Id            int    `db:"id"`
	Name          string `db:"name"`
	SubcategoryId int    `db:"subcategory_id"`
}

func (r *Characteristics) ToDomain() domains.CharacteristicsDomain {
	return domains.CharacteristicsDomain{
		Id:            r.Id,
		Name:          r.Name,
		SubcategoryId: r.SubcategoryId,
	}
}

func FromCharacteristicsDomain(domain domains.CharacteristicsDomain) Characteristics {
	return Characteristics{
		Id:            domain.Id,
		Name:          domain.Name,
		SubcategoryId: domain.SubcategoryId,
	}
}

func ToArrayOfCharacteristicsDomain(data []Characteristics) (result []domains.CharacteristicsDomain) {
	for _, v := range data {
		result = append(result, v.ToDomain())
	}
	return
}
