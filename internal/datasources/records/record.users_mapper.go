package records

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/constants"
)

func (r *Users) ToDomain() *domains.UserDomain {
	if r == nil || r.Id == 0 {
		return nil
	}
	return &domains.UserDomain{
		Id:           r.Id,
		Name:         r.Name,
		Phone:        r.Phone,
		Role:         r.Role.Name,
		CityId:       r.CityId,
		City:         *r.City.ToDomain(),
		Email:        r.Email,
		StreetNum:    r.StreetNum,
		Street:       r.Street,
		Region:       r.Region,
		Apartment:    r.Apartment,
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
		CityId:       inDom.CityId,
		Email:        inDom.Email,
		StreetNum:    inDom.StreetNum,
		Street:       inDom.Street,
		Region:       inDom.Region,
		Apartment:    inDom.Apartment,
		RefreshToken: inDom.RefreshToken,
		DateOfBirth:  inDom.DateOfBirth,
		CreatedAt:    inDom.CreatedAt,
		UpdatedAt:    inDom.UpdatedAt,
	}
}

func ToArrayOfUsersDomain(inRecs []Users) []domains.UserDomain {
	var outDom []domains.UserDomain

	for _, rec := range inRecs {
		outDom = append(outDom, *rec.ToDomain())
	}

	return outDom
}
