package seeders

import "ga_marketplace/internal/datasources/records"

var RolesData []records.Roles

func init() {
	RolesData = []records.Roles{
		{
			Id:   1,
			Name: "admin",
		},
		{
			Id:   2,
			Name: "client",
		},
	}
}
