package seeders

import (
	"ga_marketplace/internal/datasources/records"
)

var CategoriesData []records.Categories

func init() {
	CategoriesData = []records.Categories{
		{
			Id:   1,
			Name: "Category 1",
		},
		{
			Id:   2,
			Name: "Category 2",
		},
	}
}
