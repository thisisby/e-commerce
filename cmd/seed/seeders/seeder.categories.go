package seeders

import (
	"ga_marketplace/internal/datasources/records"
)

var CategoriesData []records.Categories

func init() {
	CategoriesData = []records.Categories{
		{
			Id:   1,
			Name: "No Category",
		},
		{
			Id:   2,
			Name: "Category 2",
		},
	}
}
