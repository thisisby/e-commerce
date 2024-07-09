package seeders

import "ga_marketplace/internal/datasources/records"

var CountriesData []records.Countries

func init() {
	CountriesData = []records.Countries{
		{
			Id:   1,
			Name: "Almaty",
		},
		{
			Id:   2,
			Name: "Astana",
		},
	}
}
