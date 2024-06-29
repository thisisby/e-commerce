package handlers

import (
	"ga_marketplace/internal/business/domains"
)

type RolesHandler struct {
	roleUsecase domains.RoleUsecase
}

func NewRolesHandler(roleUsecase domains.RoleUsecase) RolesHandler {
	return RolesHandler{
		roleUsecase: roleUsecase,
	}
}
