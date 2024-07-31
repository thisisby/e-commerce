package seeders

import "ga_marketplace/internal/datasources/records"

var CitiesData []records.Cities

func init() {
	CitiesData = []records.Cities{
		{
			Id:                   1,
			Name:                 "Almaty",
			DeliveryDurationDays: 3,
		},
		{
			Id:                   2,
			Name:                 "Astana",
			DeliveryDurationDays: 5,
		},
	}
}
