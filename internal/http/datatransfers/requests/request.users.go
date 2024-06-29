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

func (u *UserRegisterRequest) ToDomain() *domains.UserDomain {
	return &domains.UserDomain{
		Phone:       u.Phone,
		Name:        u.Name,
		DateOfBirth: u.DateOfBirth,
		Role:        constants.Client,
	}
}
