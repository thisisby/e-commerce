package requests

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/constants"
	"time"
)

type UserSendOTPRequest struct {
	Phone string `json:"phone" validate:"required"`
}

type UserVerifyOTPRequest struct {
	Phone string `json:"phone" validate:"required"`
	OTP   string `json:"otp" validate:"required"`
}

type UserRegisterRequest struct {
	Phone       string    `json:"phone" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	DateOfBirth time.Time `json:"date_of_birth" validate:"required"`
}

type UserUpdateRequest struct {
	CountryId   *int       `json:"country_id"`
	Name        *string    `json:"name"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	Street      *string    `json:"street"`
	Region      *string    `json:"region"`
	Apartment   *string    `json:"apartment"`
}

func (u *UserUpdateRequest) ToDomain() *domains.UserDomain {
	return &domains.UserDomain{
		CountryId: *u.CountryId,
		Name:      *u.Name,
		Street:    u.Street,
		Region:    u.Region,
		Apartment: u.Apartment,
	}

}

func (u *UserRegisterRequest) ToDomain() *domains.UserDomain {
	return &domains.UserDomain{
		Phone:       u.Phone,
		Name:        u.Name,
		DateOfBirth: u.DateOfBirth,
		Role:        constants.Client,
	}
}
