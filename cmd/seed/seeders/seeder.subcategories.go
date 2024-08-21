package seeders

import "ga_marketplace/internal/datasources/records"

var SubCategoriesData []records.SubcategoriesRecord

func init() {
	SubCategoriesData = []records.SubcategoriesRecord{
		{
			Id:         1,
			Name:       "No Category 1",
			CategoryId: 1,
		},
		{
			Id:         2,
			Name:       "SubCategory 2",
			CategoryId: 2,
		},
	}
}
