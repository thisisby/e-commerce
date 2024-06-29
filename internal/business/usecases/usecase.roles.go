package usecases

import (
	"ga_marketplace/internal/business/domains"
)

type rolesUsecase struct {
	roleRepo domains.RoleRepository
}

func NewRolesUsecase(roleRepo domains.RoleRepository) domains.RoleUsecase {
	return &rolesUsecase{
		roleRepo: roleRepo,
	}
}
