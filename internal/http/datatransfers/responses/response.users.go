package responses

import (
	"ga_marketplace/internal/business/domains"
	"time"
)

type UserResponse struct {
	Id          int           `json:"id"`
	Phone       string        `json:"phone"`
	Name        string        `json:"name"`
	DateOfBirth time.Time     `json:"date_of_birth"`
	Role        string        `json:"role"`
	CityId      int           `json:"city_id"`
	City        *CityResponse `json:"city"`
	Street      *string       `json:"street"`
	Region      *string       `json:"region"`
	Apartment   *string       `json:"apartment"`
	AccessToken string        `json:"access_token"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

func FromUserDomain(inDom *domains.UserDomain) UserResponse {
	return UserResponse{
		Id:          inDom.Id,
		Phone:       inDom.Phone,
		Name:        inDom.Name,
		DateOfBirth: inDom.DateOfBirth,
		Role:        inDom.Role,
		CityId:      inDom.CityId,
		City:        FromCityDomain(&inDom.City),
		Street:      inDom.Street,
		Region:      inDom.Region,
		Apartment:   inDom.Apartment,
		AccessToken: inDom.AccessToken,
		CreatedAt:   inDom.CreatedAt,
		UpdatedAt:   inDom.UpdatedAt,
	}
}

func FromUsersDomain(inDom []domains.UserDomain) []UserResponse {
	var outDom []UserResponse

	for _, rec := range inDom {
		outDom = append(outDom, FromUserDomain(&rec))
	}

	return outDom
}
