package seeders

import (
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
)

var ProductsData []records.Products

func init() {
	ProductsData = []records.Products{
		{
			Id:            1,
			Name:          "Products 1",
			Description:   "Description of Products 1",
			SubcategoryId: 1,
			Image:         "sad",
			Price:         100,
			Stock:         100,
			CreatedAt:     helpers.GetCurrentTime(),
			UpdatedAt:     helpers.GetCurrentTime(),
		},
		{
			Id:            2,
			Name:          "Products 2",
			Description:   "Description of Products 2",
			Image:         "sad",
			SubcategoryId: 2,
			Price:         200,
			Stock:         0,
			CreatedAt:     helpers.GetCurrentTime(),
			UpdatedAt:     helpers.GetCurrentTime(),
		},
	}
}
