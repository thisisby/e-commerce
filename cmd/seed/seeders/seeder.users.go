package seeders

import (
	"ga_marketplace/internal/constants"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
)

var UsersData []records.Users

func init() {
	UsersData = []records.Users{
		{
			Id:           1,
			Name:         "admin",
			Phone:        "+71234567890",
			RoleId:       constants.MapperRoleToId[constants.Admin],
			RefreshToken: "nil",
			DateOfBirth:  helpers.GetCurrentTime(),
			CreatedAt:    helpers.GetCurrentTime(),
			UpdatedAt:    helpers.GetCurrentTime(),
		},
		{
			Id:           2,
			Name:         "client",
			Phone:        "2",
			RoleId:       constants.MapperRoleToId[constants.Client],
			RefreshToken: "nil",
			DateOfBirth:  helpers.GetCurrentTime(),
			CreatedAt:    helpers.GetCurrentTime(),
			UpdatedAt:    helpers.GetCurrentTime(),
		},
	}
}
