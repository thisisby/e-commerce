package seeders

import "ga_marketplace/internal/datasources/records"

var BrandsData []records.Brands

func init() {
	BrandsData = []records.Brands{
		{
			Id:   1,
			Name: "No brands",
		},
		{
			Id:   2,
			Name: "Brand 2",
		},
	}
}
