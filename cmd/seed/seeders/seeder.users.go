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
			Phone:        "08123456789",
			RoleId:       constants.MapperRoleToId[constants.Admin],
			RefreshToken: "refresh_token",
			DateOfBirth:  helpers.GetCurrentTime(),
			CreatedAt:    helpers.GetCurrentTime(),
			UpdatedAt:    helpers.GetCurrentTime(),
		},
		{
			Id:           2,
			Name:         "client",
			Phone:        "123123412412",
			RoleId:       constants.MapperRoleToId[constants.Client],
			RefreshToken: "refresh_token",
			DateOfBirth:  helpers.GetCurrentTime(),
			CreatedAt:    helpers.GetCurrentTime(),
			UpdatedAt:    helpers.GetCurrentTime(),
		},
	}
}
