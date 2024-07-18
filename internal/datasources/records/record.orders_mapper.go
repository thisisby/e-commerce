package records

import "ga_marketplace/internal/business/domains"

func ToArrayOfOrdersDomain(data []Orders) []domains.OrdersDomain {
	var result []domains.OrdersDomain
	for _, val := range data {
		result = append(result, val.ToDomain())
	}

	return result
}
