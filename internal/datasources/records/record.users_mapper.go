package records

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/constants"
)

func (r *Users) ToDomain() *domains.UserDomain {
	return &domains.UserDomain{
		Id:           r.Id,
		Name:         r.Name,
		Phone:        r.Phone,
		Role:         r.Role.Name,
		RefreshToken: r.RefreshToken,
		DateOfBirth:  r.DateOfBirth,
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
	}
}

func FromUsersDomain(inDom *domains.UserDomain) Users {
	return Users{
		Id:           inDom.Id,
		Name:         inDom.Name,
		Phone:        inDom.Phone,
		RoleId:       constants.MapperRoleToId[inDom.Role],
		RefreshToken: inDom.RefreshToken,
		DateOfBirth:  inDom.DateOfBirth,
		CreatedAt:    inDom.CreatedAt,
		UpdatedAt:    inDom.UpdatedAt,
	}
}
