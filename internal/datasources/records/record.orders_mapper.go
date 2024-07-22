package records

import "ga_marketplace/internal/business/domains"

func ToArrayOfOrdersDomain(data []Orders) []domains.OrdersDomain {
	var result []domains.OrdersDomain
	for _, val := range data {
		result = append(result, val.ToDomain())
	}

	return result
}

func FromOrdersDomain(data domains.OrdersDomain) Orders {
	return Orders{
		Id:              data.Id,
		UserId:          data.UserId,
		TotalPrice:      data.TotalPrice,
		DiscountedPrice: data.DiscountedPrice,
		CityId:          data.CityId,
		City:            *FromCityDomain(&data.City),
		Status:          data.Status,
		Street:          data.Street,
		Region:          data.Region,
		Apartment:       data.Apartment,
	}
}
