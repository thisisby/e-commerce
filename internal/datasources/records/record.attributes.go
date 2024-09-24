package records

import "ga_marketplace/internal/business/domains"

type Attributes struct {
	Id                int    `db:"id"`
	Name              string `db:"name"`
	CharacteristicsId int    `db:"characteristic_id"`
}

func (r *Attributes) ToDomain() domains.AttributesDomain {
	return domains.AttributesDomain{
		Id:                r.Id,
		Name:              r.Name,
		CharacteristicsId: r.CharacteristicsId,
	}
}

func FromAttributesDomain(domain domains.AttributesDomain) Attributes {
	return Attributes{
		Id:                domain.Id,
		Name:              domain.Name,
		CharacteristicsId: domain.CharacteristicsId,
	}
}

func ToArrayOfAttributesDomain(data []Attributes) (result []domains.AttributesDomain) {
	for _, v := range data {
		result = append(result, v.ToDomain())
	}
	return
}
