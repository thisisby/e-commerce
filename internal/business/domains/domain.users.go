package domains

import (
	"time"
)

type UserDomain struct {
	Id    int
	Name  string
	Phone string
	Role  string

	CityId int
	City   CityDomain

	Street    *string
	Region    *string
	Apartment *string

	AccessToken  string
	RefreshToken string
	DateOfBirth  time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserRepository interface {
	FindByPhone(phone string) (*UserDomain, error)
	Save(inDom *UserDomain) error
	Update(inDom *UserDomain) error
	FindById(id int) (*UserDomain, error)
	FindAll() ([]UserDomain, error)
}

type UserUsecase interface {
	SendOTP(phoneNumber string) (otpCode string, statusCode int, err error)
	Save(inDom *UserDomain) (outDom *UserDomain, statusCode int, err error)
	VerifyOTP(userOTP string, redisOTP string, phone string) (statusCode int, err error)
	FindByPhone(phone string) (outDom *UserDomain, statusCode int, err error)
	Login(inDom *UserDomain) (outDom *UserDomain, statusCode int, err error)
	RefreshToken(refreshToken string) (outDom *UserDomain, statusCode int, err error)
	FindByID(id int) (outDom *UserDomain, statusCode int, err error)
	Update(inDom *UserDomain) (statusCode int, err error)
	FindAll() (outDom []UserDomain, statusCode int, err error)
}
